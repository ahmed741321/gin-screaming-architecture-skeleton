package auth

import (
	"errors"
	"screaming-architecture-skeleton/user"
)

// AuthRepository handles user authentication data access
type AuthRepository struct {
	users map[string]user.User // Simulating a database with a map
}

// NewAuthRepository initializes a new AuthRepository
func NewAuthRepository() *AuthRepository {
	return &AuthRepository{
		users: map[string]user.User{
			"user@example.com": {Email: "user@example.com", FullName: "User Example"},
		},
	}
}

// FindUserByEmail searches for a user by email
func (r *AuthRepository) FindUserByEmail(email string) (*user.User, error) {
	user, exists := r.users[email]
	if !exists {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
