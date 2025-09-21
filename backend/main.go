package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type Product struct {
	ID          int       `json:"id"`
	Category    string    `json:"category"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Date        time.Time `json:"date"`
}

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
		log.Fatalf("Faild to create connection pool: %v\n", err)
	}
	defer pool.Close()

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Get("/products", func(w http.ResponseWriter, r *http.Request) {
			rows, err := pool.Query(r.Context(), "select p.id, c.name as category, p.name, p.description, p.price, p.stock, p.created_at from products p left join categories c on p.category_id = c.id")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error fetching products: %v\n", err)
			}
			defer rows.Close()
			var products []Product

			for rows.Next() {
				var p Product
				rows.Scan(&p.ID, &p.Category, &p.Name, &p.Description, &p.Price, &p.Stock, &p.Date)
				if err != nil {
					http.Error(w, "Error iterating rows: "+err.Error(), http.StatusInternalServerError)
					return
				}
				products = append(products, p)
			}

			if err := rows.Err(); err != nil {
				http.Error(w, "Error iterating rows: "+err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(products); err != nil {
				http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
				return
			}
		})
	})

	http.ListenAndServe(":8080", r)
}
