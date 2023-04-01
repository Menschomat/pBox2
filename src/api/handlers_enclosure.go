package api

import (
	"encoding/json"
	"net/http"

	_ "github.com/Menschomat/pBox2/docs"
	m "github.com/Menschomat/pBox2/model"
)

// GetEnclosure godoc
//
//	@Summary		Get enclosure
//	@Description	Returns the whole enclosure
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
