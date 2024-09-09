package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// ENV FOR PORT
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found in env")
	}
	// Main router
	router := chi.NewRouter()

	// CORS
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	// Router for healthz
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	//.Route for error

	v1Router.Get("/err", handlerError)
	router.Mount("/v1", v1Router)

	// Server
	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server Started on PORT %v", portString)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
