package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}
type RegisterForm struct {
	Credentials
}
type LoginForm struct {
	Credentials
}
type User struct {
	ID       int
	Email    string
	Password string
	Date     time.Time
}

func AuthHandler(r chi.Router, pool *pgxpool.Pool) {
	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
		}
		var form RegisterForm
		form.Email = r.FormValue("email")
		form.Password = r.FormValue("password")

		if err := ValidateForm(form); err != nil {
			http.Error(w, "Validation error"+err.Error(), http.StatusInternalServerError)
			return
		}
		hashCost := 10
		password, err := bcrypt.GenerateFromPassword([]byte(form.Password), hashCost)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}

		if _, err := pool.Exec(r.Context(), "INSERT INTO users (email, password_hash) VALUES ($1, $2)", form.Email, password); err != nil {
			http.Error(w, "Error insert"+err.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode("User Registered"); err != nil {
			http.Error(w, "Error encode", http.StatusInternalServerError)
			return
		}
	})
	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error", http.StatusInternalServerError)
		}
		var form LoginForm
		form.Email = r.FormValue("email")
		form.Password = r.FormValue("password")
		if err := ValidateForm(form); err != nil {
			http.Error(w, "Validation error"+err.Error(), http.StatusInternalServerError)
			return
		}

		var u User
		row := pool.QueryRow(r.Context(), "SELECT * FROM users WHERE email = $1", form.Email)
		err := row.Scan(&u.ID, &u.Email, &u.Password, &u.Date)
		if err != nil {
			http.Error(w, "User not found", http.StatusInternalServerError)
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(form.Password)); err != nil {
			http.Error(w, "Password Incorrect", http.StatusInternalServerError)
			return
		}
		//Loggin
		//Create Session
		//Generate token
		//ADD EXPIRE DATE
		//INSERT TO DB
		//Send token to cookies
	})

}

func ValidateForm(form interface{}) error {
	validate := validator.New()
	if err := validate.Struct(form); err != nil {
		return err
	}
	return nil
}
