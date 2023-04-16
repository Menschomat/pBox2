package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	_ "github.com/Menschomat/pBox2/docs"
	m "github.com/Menschomat/pBox2/model"
	u "github.com/Menschomat/pBox2/utils"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-chi/chi/v5"
)

// GetFan godoc
//
//	@Summary		Get fan
//	@Description	Retrieve a fan object by its ID and the ID of the box it belongs to
//	@Tags			fan
//	@Accept			json
//	@Produce		json
//	@Param			boxId	path		int	true	"Box ID"
//	@Param			fanId	path		int	true	"Fan ID"
//	@Success		200		{object}	m.Fan
//	@Failure		400		{string}	string	"Bad Request"
//	@Router			/{boxId}/fans/{fanId} [get]
func GetFan(cfg *m.Configuration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		box := u.FindBoxById(chi.URLParam(r, "boxId"), &cfg.Enclosure)
		fan := u.FindFanById(chi.URLParam(r, "fanId"), box)
		if len(fan.ID) <= 0 {
			BadRequestError(w, r)
			return
		}
		json.NewEncoder(w).Encode(fan)
	}
}

// UpdateFan godoc
//
//	@Summary		Update fan
//	@Description	Update a fan object by its ID and the ID of the box it belongs to
//	@Tags			fan
//	@Accept			json
//	@Produce		json
//	@Param			boxId	path		int		true	"Box ID"
//	@Param			fanId	path		int		true	"Fan ID"
//	@Param			body	body		m.Fan	true	"Fan object"
//	@Success		200		{object}	m.Fan
//	@Failure		400		{string}	string	"Bad Request"
//	@Router			/{boxId}/fans/{fanId} [post]
func UpdateFan(cfg *m.Configuration, mqttClient mqtt.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		box := u.FindBoxById(chi.URLParam(r, "boxId"), &cfg.Enclosure)
		fan := u.FindFanById(chi.URLParam(r, "fanId"), box)
		if len(fan.ID) <= 0 {
			BadRequestError(w, r)
			return
		}

		var body m.Fan
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			BadRequestError(w, r)
			return
		}

		fan.Level = body.Level
		mqttClient.Publish("test/"+box.ID+"/fans/"+fan.ID, 0, false, strconv.Itoa(fan.Level))
		json.NewEncoder(w).Encode(fan)
	}
}
