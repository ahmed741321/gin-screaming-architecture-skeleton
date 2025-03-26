package user

import "errors"

// User represents a user model
type User struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Password string `json:"-"` // Add password field for authentication
}

// UserRepository handles data access for users
type UserRepository struct {
	users map[string]User // Simulating a database
}

// NewUserRepository initializes a new UserRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: map[string]User{
			"jane@example.com": {Email: "jane@example.com", FullName: "Jane Doe", Password: "password123"},
			"john@example.com": {Email: "john@example.com", FullName: "John Smith", Password: "password123"},
		},
	}
}

// FindUserByEmail searches for a user by email
func (r *UserRepository) FindUserByEmail(email string) (*User, error) {
	user, exists := r.users[email]
	if !exists {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
