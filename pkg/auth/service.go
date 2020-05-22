package auth

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrInvalidCredentials ...
	ErrInvalidCredentials = errors.New("Invalid username or password")
)

// Service ...
type Service interface {
	Login(string, string) (*User, error)
	Logout() error
	Register(string, string) (*User, error)
}

// Repository ...
type Repository interface {
	GetUser(string) (*User, error)
	CreateUser(*User) error
	// Logout() error
}

type service struct {
	r Repository
}

// NewService creates an auth service with the necessary dependencies.
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Login(email, pwd string) (*User, error) {
	user, err := s.r.GetUser(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(pwd))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, err
}

func (s *service) Logout() error {
	// err := s.r.Logout()
	return nil
}

func (s *service) Register(email, pwd string) (*User, error) {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(pwd), 11)
	if err != nil {
		return nil, errors.New("Failed to register user")
	}

	u, err := s.r.GetUser(email)
	if u != nil {
		return nil, errors.New("User already exists")
	}

	user := User{Email: email, PasswordHash: string(pwdHash), Name: strings.Split(email, "@")[0]}
	if err = s.r.CreateUser(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
