package auth

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AuthHandler(r chi.Router, pool *pgxpool.Pool) {
	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode("register"); err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
		}
	})
}

