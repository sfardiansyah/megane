package mongodb

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/sfardiansyah/megane/pkg/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Storage ...
type Storage struct {
	db *mongo.Database
}

// NewStorage ...
func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	c, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if err != nil {
		return nil, err
	}

	if err := c.Connect(context.TODO()); err != nil {
		return nil, errors.New("Failed to connect to the database")
	}

	s.db = c.Database(os.Getenv("MONGODB_DB"))

	return s, nil
}

// GetUser ...
func (s *Storage) GetUser(email string) (*auth.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user := new(auth.User)

	collection := s.db.Collection("user")

	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, errors.New("Invalid username or password")
	}

	return user, nil
}

// CreateUser ...
func (s *Storage) CreateUser(user *auth.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := s.db.Collection("user")

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return errors.New("Failed to register user")
	}

	return nil
}
