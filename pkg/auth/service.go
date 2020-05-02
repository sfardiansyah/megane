package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Service ...
type Service interface {
	Login(string, string) (User, error)
	Logout() error
	Register() (User, error)
}

// Repository ...
type Repository interface {
	GetUser(string) (User, error)
	// Logout() error
	// Register() (User, error)
}

type service struct {
	r Repository
}

// NewService creates an auth service with the necessary dependencies.
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Login(email, password string) (User, error) {
	user, err := s.r.GetUser(email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, errors.New("Invalid username or password")
	}

	return user, err
}

func (s *service) Logout() error {
	// err := s.r.Logout()
	return nil
}

func (s *service) Register() (User, error) {
	// user, err := s.r.Register()
	return User{}, nil
}
