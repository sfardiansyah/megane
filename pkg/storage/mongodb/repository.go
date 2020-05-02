package mongodb

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/sfardiansyah/megane/pkg/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Storage ...
type Storage struct {
	db *mongo.Client
}

// NewStorage ...
func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	s.db, err = mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return s, nil
}

// GetUser ...
func (s *Storage) GetUser(email string) (auth.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user := auth.User{}

	err := s.db.Connect(ctx)
	if err != nil {
		return user, errors.New("Failed to connect to the database")
	}

	collection := s.db.Database(os.Getenv("MONGODB_DB")).Collection("user")

	err = collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return user, errors.New("Invalid username or password")
	}

	return user, nil
}
