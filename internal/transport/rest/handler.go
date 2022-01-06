package rest

import (
	"github.com/gorilla/mux"
	"github.com/p12s/using-orm-with-db/internal/service"
	"net/http"
)

// Handler
type Handler struct {
	services *service.Service
}

// NewHandler - constructor
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// InitRouter - init routes
func (h *Handler) InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)
	router.HandleFunc("/health", h.health).Methods(http.MethodGet)
	router.HandleFunc("/sign-up", h.signUp).Methods(http.MethodPost)
	router.HandleFunc("/sign-in", h.signIn).Methods(http.MethodPost)

	account := router.PathPrefix("/account").Subrouter()
	{
		account.Use(h.authMiddleware)
		account.HandleFunc("/", h.getAccount).Methods(http.MethodGet)
	}

	product := router.PathPrefix("/product").Subrouter()
	{
		product.Use(h.authMiddleware)
		product.HandleFunc("/", h.health).Methods(http.MethodPost)              // TODO
		product.HandleFunc("/{id:[0-9]+}", h.health).Methods(http.MethodPut)    // TODO
		product.HandleFunc("/{id:[0-9]+}", h.health).Methods(http.MethodDelete) // TODO
		product.HandleFunc("/", h.health).Methods(http.MethodGet)               // TODO
	}

	return router
}
