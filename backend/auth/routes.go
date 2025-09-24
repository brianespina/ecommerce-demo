package auth

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type RegisterForm struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func AuthHandler(r chi.Router, pool *pgxpool.Pool) {
	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
		}
		form := RegisterForm{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}
		validate := validator.New()

		hashCost := 10
		password, err := bcrypt.GenerateFromPassword([]byte(form.Password), hashCost)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
		}

		if err := validate.Struct(form); err != nil {
			http.Error(w, "Validation error"+err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := pool.Exec(r.Context(), "INSERT INTO users (email, password_hash) VALUES ($1, $2)", form.Email, password); err != nil {
			http.Error(w, "Error insert"+err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode([]string{form.Email, form.Password}); err != nil {
			http.Error(w, "Error encode", http.StatusInternalServerError)
			return
		}
	})
	r.Post("login", func(w http.ResponseWriter, r *http.Request) {

	})
}
