package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CategoriesHandler(r chi.Router, pool *pgxpool.Pool) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		rows, err := pool.Query(r.Context(), "SELECT * FROM categories")
		if err != nil {
			http.Error(w, "Error fetching categories: %v\n", http.StatusInternalServerError)
		}
		defer rows.Close()
		var categories []Category
		for rows.Next() {
			var category Category
			err := rows.Scan(&category.ID, &category.Name)
			if err != nil {
				http.Error(w, "Error Scanning Rows", http.StatusInternalServerError)
				return
			}
			categories = append(categories, category)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, "Error Iterating rows", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(categories); err != nil {
			http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
