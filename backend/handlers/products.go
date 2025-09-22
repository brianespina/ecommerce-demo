package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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

func ProductsHandler(r chi.Router, pool *pgxpool.Pool) {
	//Get products by ID
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		rawParam := chi.URLParam(r, "id")
		id, err := strconv.Atoi(rawParam)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusNotFound)
			return
		}
		var p Product
		if err := pool.QueryRow(r.Context(), "SELECT * FROM products WHERE id = $1", id).Scan(&p.ID, &p.Category, &p.Name, &p.Description, &p.Price, &p.Stock, &p.Date); err != nil {
			http.Error(w, "Product not found: "+err.Error(), http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(w).Encode(p); err != nil {
			http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
			return
		}
	})

	//Get all products
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		rows, err := pool.Query(r.Context(), "select p.id, c.name as category, p.name, p.description, p.price, p.stock, p.created_at from products p left join categories c on p.category_id = c.id")
		if err != nil {
			http.Error(w, "Error fetching products: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		var products []Product

		for rows.Next() {
			var p Product
			err := rows.Scan(&p.ID, &p.Category, &p.Name, &p.Description, &p.Price, &p.Stock, &p.Date)
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
}
