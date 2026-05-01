package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func AllowOriginFunc(r *http.Request, origin string) bool {
	// TODO: []string{"http://*", "https://*"}
	return true
}

func main() {
	println("Hello, World!")

	godotenv.Load() // Load environment variables from .env file
	portStr := os.Getenv("PORT")
	if portStr == "" {
		println("PORT environment variable is not set.")
	} else {
		println("PORT environment variable is set to:", portStr)
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowOriginFunc:  AllowOriginFunc,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portStr,
	}

	log.Printf("Starting server on port %s...\n", portStr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
