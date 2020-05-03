package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sfardiansyah/megane/pkg/auth"
)

// Handler ...
func Handler(a auth.Service) http.Handler {
	r := mux.NewRouter()

	s := r.PathPrefix("/api/v1").Subrouter()

	s.HandleFunc("/login", login(a)).Methods("POST")
	s.HandleFunc("/register", register(a)).Methods("POST")

	return r
}

func login(a auth.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		lR := LoginRequest{}
		if err := json.NewDecoder(r.Body).Decode(&lR); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := a.Login(lR.Email, lR.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func register(a auth.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rR := RegisterRequest{}
		if err := json.NewDecoder(r.Body).Decode(&rR); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := a.Register(rR.Email, rR.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
