package rest

import (
	"encoding/json"
	"net/http"

	_ "github.com/Menschomat/pBox2/docs"
	m "github.com/Menschomat/pBox2/model"
	u "github.com/Menschomat/pBox2/utils"
	"github.com/go-chi/chi/v5"
)

// GetSensor godoc
//
//	@Summary		Get Sensor
//	@Description	Get a sensor by its box and sensor ID.
//	@Tags			sensor
//	@Accept			json
//	@Produce		json
//	@Param			boxId		path		string	true	"Box ID"
//	@Param			sensorId	path		string	true	"Sensor ID"
//	@Success		200			{object}	m.Sensor
//	@Failure		400			{string}	string	"Bad Request"
//	@Router			/{boxId}/sensors/{sensorId} [get]
func GetSensor(cfg *m.Configuration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		box := u.FindBoxById(chi.URLParam(r, "boxId"), &cfg.Enclosure)
		sensor := u.FindSensorById(chi.URLParam(r, "sensorId"), box)
		if len(sensor.ID) <= 0 {
			BadRequestError(w, r)
			return
		}
		json.NewEncoder(w).Encode(sensor)
	}
}

// GetSensorData godoc
//
//	@Summary		Get Sensor Data
//	@Description	Get time-series data for a sensor by its box and sensor ID.
//	@Tags			sensor
//	@Accept			json
//	@Produce		json
//	@Param			boxId		path		string	true	"Box ID"
//	@Param			sensorId	path		string	true	"Sensor ID"
//	@Success		200			{object}	m.TimeSeries
//	@Failure		400			{string}	string	"Bad Request"
//	@Router			/{boxId}/sensors/{sensorId}/data [get]
func GetSensorData(cfg *m.Configuration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		box := u.FindBoxById(chi.URLParam(r, "boxId"), &cfg.Enclosure)
		sensor := u.FindSensorById(chi.URLParam(r, "sensorId"), box)
		if len(sensor.ID) <= 0 {
			BadRequestError(w, r)
			return
		}
		json.NewEncoder(w).Encode(sensor.TimeSeries)
	}
}

// @Summary		Bad Request Error
// @Description	Returns a 400 Bad Request error response
// @Tags			error
// @Failure		400	{string}	string	"Bad Request"
func BadRequestError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("bad request error"))
}
