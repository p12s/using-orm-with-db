package rest

import (
	"fmt"
	"net/http"
	"os"
)

// ErrorResponse - when not Ok
func ErrorResponse(w http.ResponseWriter, statusCode int, userMessage, logMessage string) {
	_, _ = fmt.Fprintf(os.Stderr, "fail occurred: %s", logMessage)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write([]byte(userMessage))
}

// OkResponse - when Ok
func OkResponse(w http.ResponseWriter, userMessage []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(userMessage)
}
