package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Menschomat/pBox2/api/websocket"
	_ "github.com/Menschomat/pBox2/docs"
	m "github.com/Menschomat/pBox2/model"
	u "github.com/Menschomat/pBox2/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-chi/chi/v5"
)

// respondWithError writes a json response with an error message and a status code.
func respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// respondWithJSON writes a json response with a payload and a status code.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

// GetFan handles the retrieval of a fan by its ID and the ID of its parent box.
// @Summary Retrieve a fan by its ID.
// @Description Get the details of a specific fan by its unique ID and the ID of the box it belongs to.
// @Tags fan
// @Accept json
// @Produce json
// @Param boxId path int true "Box ID"
// @Param fanId path int true "Fan ID"
// @Success 200 {object} m.Fan "Successfully retrieved fan"
// @Failure 400 {object} string "Invalid ID format or fan not found"
// @Router /{boxId}/fans/{fanId} [get]
func GetFan(cfg *m.Configuration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		boxID := chi.URLParam(r, "boxId")
		fanID := chi.URLParam(r, "fanId")

		box := u.FindBoxById(boxID, &cfg.Enclosure)
		fan := u.FindFanById(fanID, box)
		if fan == nil || len(fan.ID) == 0 {
			respondWithError(w, http.StatusBadRequest, "Invalid box ID or fan ID")
			return
		}

		respondWithJSON(w, http.StatusOK, fan)
	}
}

// UpdateFan handles the update of a fan's information.
// @Summary Update a fan's information.
// @Description Update the information of a specific fan by its unique ID and the ID of the box it belongs to.
// @Tags fan
// @Accept json
// @Produce json
// @Param boxId path int true "Box ID"
// @Param fanId path int true "Fan ID"
// @Param body body m.Fan true "Fan object to update"
// @Success 200 {object} m.Fan "Successfully updated fan"
// @Failure 400 {object} string "Invalid payload format or fan not found"
// @Router /{boxId}/fans/{fanId} [post]
func UpdateFan(cfg *m.Configuration, mqttClient mqtt.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		boxID := chi.URLParam(r, "boxId")
		fanID := chi.URLParam(r, "fanId")

		box := u.FindBoxById(boxID, &cfg.Enclosure)
		fan := u.FindFanById(fanID, box)
		if fan == nil || len(fan.ID) == 0 {
			respondWithError(w, http.StatusBadRequest, "Invalid box ID or fan ID")
			return
		}

		var updatedFan m.Fan
		if err := json.NewDecoder(r.Body).Decode(&updatedFan); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		fan.Level = updatedFan.Level

		// Validate or process the new fan level before publishing it.
		topic := "test/" + box.ID + "/fans/" + fan.ID
		mqttClient.Publish(topic, 0, false, strconv.Itoa(fan.Level))
		websocket.PublishFanEvent(cfg, box, fan)
		respondWithJSON(w, http.StatusOK, fan)
	}
}
