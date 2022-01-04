package rest

import (
	"encoding/json"
	"net/http"
)

// health - app technical service handler
func (h *Handler) health(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(map[string]string{
		"service": "orm-training",
		"status":  "Ok",
	})

	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "internal error", err.Error())
		return
	}

	OkResponse(w, response)
}
