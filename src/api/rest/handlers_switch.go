package rest

import (
	"encoding/json"
	"net/http"

	api_mqtt "github.com/Menschomat/pBox2/api/mqtt"
	"github.com/Menschomat/pBox2/api/websocket"
	_ "github.com/Menschomat/pBox2/docs"
	m "github.com/Menschomat/pBox2/model"
	u "github.com/Menschomat/pBox2/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-chi/chi/v5"
)

// GetSwitch handles the retrieval of a switch by its ID and the ID of its parent box.
// @Summary Retrieve a switch by its ID.
// @Description Get the details of a specific switch by its unique ID and the ID of the box it belongs to.
// @Tags switch
// @Accept json
// @Produce json
// @Param boxId path		string	true "Box ID"
// @Param switchId path		string	true "Switch ID"
// @Success 200 {object} m.Switch "Successfully retrieved switch"
// @Failure 400 {object} string "Invalid ID format or switch not found"
// @Router /{boxId}/switches/{switchId} [get]
func GetSwitch(cfg *m.Configuration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		boxID := chi.URLParam(r, "boxId")
		switchID := chi.URLParam(r, "switchId")

		box := u.FindBoxById(boxID, &cfg.Enclosure)
		switc := u.FindSwitchById(switchID, box)
		if switc == nil || len(switc.ID) == 0 {
			respondWithError(w, http.StatusBadRequest, "Invalid box ID or switch ID")
			return
		}

		respondWithJSON(w, http.StatusOK, switc)
	}
}

// UpdateSwitch handles the update of a switch's state.
// @Summary Update a switch's state.
// @Description Update the state of a specific switch by its unique ID and the ID of the box it belongs to.
// @Tags switch
// @Accept json
// @Produce json
// @Param boxId path		string	true "Box ID"
// @Param switchId path		string	true "Switch ID"
// @Param body body m.Switch true "Switch object with new state"
// @Success 200 {object} m.Switch "Successfully updated switch"
// @Failure 400 {object} string "Invalid payload format or switch not found"
// @Router /{boxId}/switches/{switchId} [post]
func UpdateSwitch(cfg *m.Configuration, mqttClient mqtt.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		boxID := chi.URLParam(r, "boxId")
		switchID := chi.URLParam(r, "switchId")

		box := u.FindBoxById(boxID, &cfg.Enclosure)
		switc := u.FindSwitchById(switchID, box)
		if switc == nil || len(switc.ID) == 0 {
			respondWithError(w, http.StatusBadRequest, "Invalid box ID or switch ID")
			return
		}

		var updatedSwitch m.Switch
		if err := json.NewDecoder(r.Body).Decode(&updatedSwitch); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()

		switc.State = updatedSwitch.State

		// Publish the new switch state to the MQTT topic.
		topic := cfg.Mqtt.Tasmota.Topic + "/cmnd/" + switchID + "/POWER"
		stateStr := api_mqtt.BoolToStr(switc.State)
		mqttClient.Publish(topic, 0, false, stateStr)
		websocket.PublishSwitchEvent(cfg, box, switc)
		respondWithJSON(w, http.StatusOK, switc)
	}
}
