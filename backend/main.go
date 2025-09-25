package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"espinabrian.com/ecommerce/auth"
	"espinabrian.com/ecommerce/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConfig := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	pool, err := pgxpool.New(context.Background(), dbConfig)

	if err != nil {
		log.Fatalf("Failed to create connection pool: %v\n", err)
	}
	defer pool.Close()

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/api", func(r chi.Router) {

		r.Route("/auth", func(r chi.Router) {
			auth.AuthHandler(r, pool)
		})

		r.Route("/products", func(r chi.Router) {
			handlers.ProductsHandler(r, pool)
		})

		r.Route("/categories", func(r chi.Router) {
			handlers.CategoriesHandler(r, pool)
		})
	})

	http.ListenAndServe(":8080", r)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("session_token")
		if err != nil {
			http.Error(w, "Login required", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
