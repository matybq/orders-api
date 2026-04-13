package app

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/matybq/orders-api/handler"
)
// loadRoutes: A function that sets up the routing for the application using the chi router. 
// It defines all the routes and their corresponding handlers.
func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", homeHandler)

	router.Route("/orders", loadOrderRoutes)

	return router
}

// ResponseWriter: An interface used to construct an HTTP response.
// Request: A struct that represents an HTTP request received by the server.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the Orders API!"))
}

func loadOrderRoutes(router chi.Router) {
	orderHandler := &handler.Order{}

	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.Get)
	router.Put("/{id}", orderHandler.Update)
	router.Delete("/{id}", orderHandler.Delete)
}
