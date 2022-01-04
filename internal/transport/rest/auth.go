package rest

/*
func (h *rest.Handler) signUp(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		rest.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var input domain.SignUpInput
	if err = json.Unmarshal(reqBytes, &input); err != nil {
		rest.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = input.Validate(); err != nil {
		rest.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.CreateUser(r.Context(), input)
	if err != nil {
		rest.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *rest.Handler) signIn(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		rest.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var input domain.SignInInput
	if err = json.Unmarshal(reqBytes, &input); err != nil {
		rest.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = input.Validate(); err != nil {
		rest.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.GetUserByCredentials(r.Context(), input.Email, input.Password)
	if err != nil {
		rest.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response, err := json.Marshal(map[string]string{
		"token": token,
	})
	if err != nil {
		rest.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}
*/
