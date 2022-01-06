package rest

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) getAccount(w http.ResponseWriter, r *http.Request) {
	accountId, err := getAccountId(r)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "unknown account", err.Error())
		return
	}

	account, err := h.services.GetAccountById(accountId)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "service failure", err.Error())
		return
	}

	response, err := json.Marshal(account)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "service failure", err.Error())
		return
	}

	OkResponse(w, response)
}
