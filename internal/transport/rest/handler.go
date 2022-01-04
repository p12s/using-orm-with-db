package rest

import (
	"github.com/gorilla/mux"
	"net/http"

	"github.com/p12s/using-orm-with-db/internal/service"
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

	//auth := r.PathPrefix("/products").Subrouter()
	//{
	//	products.HandleFunc("/", h.createProduct).Methods(http.MethodPost)
	//	products.HandleFunc("/{id:[0-9]+}", h.updateProduct).Methods(http.MethodPut)
	//	products.HandleFunc("/{id:[0-9]+}", h.deleteProduct).Methods(http.MethodDelete)
	//	products.HandleFunc("/", h.getAllProducts).Methods(http.MethodGet)
	//}

	return router
} // TODO как сделать инит разных групп роутов - просто в разных папках?
