package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/prantoran/rssagg/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func AllowOriginFunc(r *http.Request, origin string) bool {
	// TODO: []string{"http://*", "https://*"}
	return true
}

func main() {
	godotenv.Load() // Load environment variables from .env file
	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("PORT environment variable is not set.")
	} else {
		println("PORT environment variable is set to:", portStr)
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set.")
	} else {
		println("DB_URL environment variable is set to:", dbURL)
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer conn.Close()

	queries := database.New(conn)

	apiCfg := &apiConfig{
		DB: queries,
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
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.authMiddleware(apiCfg.handlerGetUser))
	v1Router.Post("/feeds", apiCfg.authMiddleware(apiCfg.handlerCreateFeed))
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
