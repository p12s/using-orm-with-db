package rest

import (
	"encoding/json"
	"github.com/p12s/using-orm-with-db/internal/domain"
	"io"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "error reading post-data", err.Error())
		return
	}

	var input domain.SignUpInput
	if err = json.Unmarshal(reqBytes, &input); err != nil {
		ErrorResponse(w, http.StatusBadRequest, "invalid input body", err.Error())
		return
	}

	if err = input.Validate(); err != nil {
		ErrorResponse(w, http.StatusBadRequest, "invalid input body", err.Error())
		return
	}

	err = h.services.CreateAccount(input)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "service failure", err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, "error reading post-data", err.Error())
		return
	}

	var input domain.SignInInput
	if err = json.Unmarshal(reqBytes, &input); err != nil {
		ErrorResponse(w, http.StatusBadRequest, "invalid input body", err.Error())
		return
	}

	// TODO maybe return more specific error texts - what is wrong (email ...)
	if err = input.Validate(); err != nil {
		ErrorResponse(w, http.StatusBadRequest, "invalid input body", err.Error())
		return
	}

	token, err := h.services.GetTokenByCredentials(input)
	if err != nil {
		ErrorResponse(w, http.StatusForbidden, "invalid input body", err.Error())
		return
	}

	response, err := json.Marshal(map[string]string{
		"token": token,
	})
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, "service failure", err.Error())
		return
	}

	OkResponse(w, response)
}
