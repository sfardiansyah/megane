package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sfardiansyah/megane/pkg/auth"
)

// Handler ...
func Handler(a auth.Service) http.Handler {
	r := mux.NewRouter()

	s := r.PathPrefix("/api/v1").Subrouter()

	s.HandleFunc("/login", login(a)).Methods("POST")
	s.HandleFunc("/register", register(a)).Methods("POST")

	return handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(r)
}

func login(a auth.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		lR := LoginRequest{}
		if err := json.NewDecoder(r.Body).Decode(&lR); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, token, err := a.Login(lR.Email, lR.Password)
		if err != nil {
			if errors.Is(err, auth.ErrInvalidCredentials) {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		payload := map[string]interface{}{
			"user":  user,
			"token": token,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(payload)
	}
}

func register(a auth.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rR := RegisterRequest{}
		if err := json.NewDecoder(r.Body).Decode(&rR); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, token, err := a.Register(rR.Email, rR.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		payload := map[string]interface{}{
			"user":  user,
			"token": token,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(payload)
	}
}
