package rest

import (
	"bytes"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/p12s/using-orm-with-db/internal/domain"
	"github.com/p12s/using-orm-with-db/internal/service"
	mock_service "github.com/p12s/using-orm-with-db/internal/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_signUp(t *testing.T) {
	type authMockBehavior func(s *mock_service.MockAuther, input domain.SignUpInput)

	tests := []struct {
		name                string
		inputBody           string
		inputSignUp         domain.SignUpInput
		authMockBehavior    authMockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "Can sign up with correct input",
			inputBody: `{"password": "qwerty", "email": "test@test.ru"}`,
			inputSignUp: domain.SignUpInput{
				Password: "qwerty",
				Email:    "test@test.ru",
			},
			authMockBehavior: func(s *mock_service.MockAuther, input domain.SignUpInput) {
				s.EXPECT().CreateAccount(input).Return(nil)
			},
			expectedStatusCode:  http.StatusCreated,
			expectedRequestBody: ``,
		},
		{
			name:      "Can't sign up with input without email",
			inputBody: `{"password": "qwerty"}`,
			inputSignUp: domain.SignUpInput{
				Password: "qwerty",
			},
			authMockBehavior:    func(s *mock_service.MockAuther, input domain.SignUpInput) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Can't sign up with input without password",
			inputBody: `{"email": "test@test.ru"}`,
			inputSignUp: domain.SignUpInput{
				Email: "test@test.ru",
			},
			authMockBehavior:    func(s *mock_service.MockAuther, input domain.SignUpInput) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Can return error response if service failure",
			inputBody: `{"password": "qwerty", "email": "test@test.ru"}`,
			inputSignUp: domain.SignUpInput{
				Password: "qwerty",
				Email:    "test@test.ru",
			},
			authMockBehavior: func(s *mock_service.MockAuther, input domain.SignUpInput) {
				s.EXPECT().CreateAccount(input).Return(errors.New(""))
			},
			expectedStatusCode:  http.StatusInternalServerError,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			acc := mock_service.NewMockAuther(ctrl)
			tt.authMockBehavior(acc, tt.inputSignUp)
			serviceMock := &service.Service{Auther: acc}

			h := NewHandler(serviceMock)
			r := mux.NewRouter()
			r.HandleFunc("/sign-up", h.signUp).Methods(http.MethodPost)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(tt.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_signIn(t *testing.T) {
	type authMockBehavior func(s *mock_service.MockAuther, input domain.SignInInput)

	tests := []struct {
		name                string
		inputBody           string
		inputSignIn         domain.SignInInput
		authMockBehavior    authMockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "Can sign in with correct input",
			inputBody: `{"password": "qwerty", "email": "test@test.ru"}`,
			inputSignIn: domain.SignInInput{
				Password: "qwerty",
				Email:    "test@test.ru",
			},
			authMockBehavior: func(s *mock_service.MockAuther, input domain.SignInInput) {
				s.EXPECT().GetTokenByCredentials(input).Return("token", nil)
			},
			expectedStatusCode:  http.StatusOK,
			expectedRequestBody: `{"token":"token"}`,
		},
		{
			name:      "Can't sign in with input without email",
			inputBody: `{"password": "qwerty"}`,
			inputSignIn: domain.SignInInput{
				Password: "qwerty",
			},
			authMockBehavior:    func(s *mock_service.MockAuther, input domain.SignInInput) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Can't sign in with input without password",
			inputBody: `{"email": "test@test.ru"}`,
			inputSignIn: domain.SignInInput{
				Email: "test@test.ru",
			},
			authMockBehavior:    func(s *mock_service.MockAuther, input domain.SignInInput) {},
			expectedStatusCode:  http.StatusBadRequest,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Can return error response if service failure",
			inputBody: `{"password": "qwerty", "email": "test@test.ru"}`,
			inputSignIn: domain.SignInInput{
				Password: "qwerty",
				Email:    "test@test.ru",
			},
			authMockBehavior: func(s *mock_service.MockAuther, input domain.SignInInput) {
				s.EXPECT().GetTokenByCredentials(input).Return("token", errors.New(""))
			},
			expectedStatusCode:  http.StatusForbidden,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			acc := mock_service.NewMockAuther(ctrl)
			tt.authMockBehavior(acc, tt.inputSignIn)
			serviceMock := &service.Service{Auther: acc}

			h := NewHandler(serviceMock)
			r := mux.NewRouter()
			r.HandleFunc("/sign-in", h.signIn).Methods(http.MethodPost)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-in", bytes.NewBufferString(tt.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedRequestBody, w.Body.String())
		})
	}
}
