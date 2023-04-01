package api

import (
	"encoding/json"
	"net/http"

	_ "github.com/Menschomat/pBox2/docs"
	m "github.com/Menschomat/pBox2/model"
	u "github.com/Menschomat/pBox2/utils"
	"github.com/go-chi/chi/v5"
)

// GetEnclosure godoc
//
//	@Summary		returns whole enclosure
//	@Description	get string by ID
//	@Tags			enclosure
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	m.Enclosure
//	@Router			/enclosure [get]
func GetEnclosure(cfg *m.Configuration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(cfg.Enclosure)
	}
}

// UpdateFan godoc
//
//	@Summary		updates fan
//	@Description	get fan by box- and fan-id
//	@Tags			fan
//	@Accept			json
//	@Produce		json
//	@Param			request	body		m.Fan	true	"body"
//	@Success		200		{object}	m.Fan
//	@Router			/{boxId}/fans/{fanId} [post]
func UpdateFan(cfg *m.Configuration) http.HandlerFunc {
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
		json.NewEncoder(w).Encode(fan)
	}
}

// UpdateLight godoc
//
//	@Summary		updates light
//	@Description	get light by box- and light-id
//	@Tags			light
//	@Accept			json
//	@Produce		json
//	@Param			request	body		m.Light	true	"body"
//	@Success		200		{object}	m.Light
//	@Router			/{boxId}/lights/{lightId} [post]
func UpdateLight(cfg *m.Configuration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		box := u.FindBoxById(chi.URLParam(r, "boxId"), &cfg.Enclosure)
		light := u.FindLightById(chi.URLParam(r, "lightId"), box)
		if len(light.ID) <= 0 {
			BadRequestError(w, r)
			return
		}

		var body m.Light
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			BadRequestError(w, r)
			return
		}

		light.Level = body.Level
		json.NewEncoder(w).Encode(light)
	}
}

// GetFan godoc
//
//	@Summary		returns fan
//	@Description	get fan by box- and fan-id
//	@Tags			fan
//	@Accept			json
//	@Produce		json
//	@Param			boxId	path		string	true	"Box ID"
//	@Param			fanId	path		string	true	"Fan ID"
//	@Success		200		{object}	m.Fan
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

// GetLight godoc
//
//	@Summary		returns light
//	@Description	get light by box- and light-id
//	@Tags			light
//	@Accept			json
//	@Produce		json
//	@Param			boxId	path		string	true	"Box ID"
//	@Param			lightId	path		string	true	"Light ID"
//	@Success		200		{object}	m.Light
//	@Router			/{boxId}/lights/{lightId} [get]
func GetLight(cfg *m.Configuration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		box := u.FindBoxById(chi.URLParam(r, "boxId"), &cfg.Enclosure)
		light := u.FindLightById(chi.URLParam(r, "lightId"), box)
		if len(light.ID) <= 0 {
			BadRequestError(w, r)
			return
		}
		json.NewEncoder(w).Encode(light)
	}
}

// GetSensor godoc
//
//	@Summary		returns sensor
//	@Description	get sensor by box- and sensor-id
//	@Tags			sensor
//	@Accept			json
//	@Produce		json
//	@Param			boxId		path		string	true	"Box ID"
//	@Param			sensorId	path		string	true	"Sensor ID"
//	@Success		200			{object}	m.Sensor
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
//	@Summary		returns sensor-data as time-series
//	@Description	get sensor-data as time-series by box- and sensor-id
//	@Tags			sensor
//	@Accept			json
//	@Produce		json
//	@Param			boxId		path		string	true	"Box ID"
//	@Param			sensorId	path		string	true	"Sensor ID"
//	@Success		200			{object}	m.TimeSeries
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

func BadRequestError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("bad request error"))
}
