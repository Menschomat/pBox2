package api

import (
	"encoding/json"
	"net/http"

	_ "github.com/Menschomat/pBox2/docs"
	m "github.com/Menschomat/pBox2/model"
	u "github.com/Menschomat/pBox2/utils"
	"github.com/go-chi/chi/v5"
)

// GetLight godoc
//
//	@Summary		Get light
//	@Description	Get information about a specific light in a specific box
//	@Tags			light
//	@Accept			json
//	@Produce		json
//	@Param			boxId	path		string	true	"Box ID"
//	@Param			lightId	path		string	true	"Light ID"
//	@Success		200		{object}	m.Light
//	@Failure		400		{string}	string	"Bad Request"
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

// UpdateLight godoc
//
//	@Summary		Update light
//	@Description	Update the level of a specific light in a specific box
//	@Tags			light
//	@Accept			json
//	@Produce		json
//	@Param			boxId	path		string	true	"Box ID"
//	@Param			lightId	path		string	true	"Light ID"
//	@Param			request	body		m.Light	true	"Light object"
//	@Success		200		{object}	m.Light
//	@Failure		400		{string}	string	"Bad Request"
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
