package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/hello", basicHandler)

	server := &http.Server{
		Addr: ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	// error handling: if the server fails to start, it will print an error message to the console.
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// ResponseWriter: An interface used to construct an HTTP response.
// Request: A struct that represents an HTTP request received by the server.
func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}
