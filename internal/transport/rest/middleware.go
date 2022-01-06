package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// CtxValue
type CtxValue int

const (
	ctxUserID CtxValue = iota
)

// loggingMiddleware - logging
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// logging-example, if it need:
		// _, _ = fmt.Fprintf(os.Stdout, "%s: [%s] - %s ", time.Now().Format(time.RFC3339), r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

// authMiddleware - authentication
func (h *Handler) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := getTokenFromRequest(r)
		if err != nil {
			ErrorResponse(w, http.StatusUnauthorized, "empty or wrong token", err.Error())
			return
		}

		userId, err := h.services.ParseToken(token)
		if err != nil {
			ErrorResponse(w, http.StatusUnauthorized, "invalid token", err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), ctxUserID, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// getTokenFromRequest
func getTokenFromRequest(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return headerParts[1], nil
}

// getUserId
func getAccountId(r *http.Request) (int, error) {
	accId := r.Context().Value(ctxUserID)

	accIdAsInt, ok := accId.(int)
	if !ok {
		return 0, fmt.Errorf("undefined user id: %v", accIdAsInt)
	}
	if accIdAsInt <= 0 {
		return 0, fmt.Errorf("negative user id: %v", accIdAsInt)
	}

	return accIdAsInt, nil
}
