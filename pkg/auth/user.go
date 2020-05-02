package auth

// User defines the properties of a user account
type User struct {
	Email        string `json:"email" bson:"email"`
	PasswordHash string `json:"-" bson:"passwordHash"`
	Name         string `json:"name" bson:"name"`
}
