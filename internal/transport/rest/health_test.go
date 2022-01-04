package rest

import (
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/p12s/using-orm-with-db/internal/service"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_health(t *testing.T) {
	tests := []struct {
		name                string
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:                "Can return service name and status",
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `{"service":"orm-training","status":"Ok"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// controller mock
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// service mock
			serviceMock := &service.Service{Auther: nil}

			// handler mock
			h := NewHandler(serviceMock)
			r := mux.NewRouter()
			r.HandleFunc("/health", h.health).Methods(http.MethodGet)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/health", nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedRequestBody, w.Body.String())
		})
	}
}
