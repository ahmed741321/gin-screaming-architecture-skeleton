package auth

import (
	"errors"
	"screaming-architecture-skeleton/user"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateJWT(user user.User) string {
	// Placeholder for JWT generation logic
	return "mocked-jwt-token"
}

type AuthService struct {
	authRepo *AuthRepository // Change to pointer
}

func NewAuthService(authRepo *AuthRepository) *AuthService { // Accept pointer
	return &AuthService{authRepo: authRepo}
}

func (s *AuthService) Authenticate(credentials LoginRequest) (string, error) {
	user, err := s.authRepo.FindUserByEmail(credentials.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Simulate password check (password field is missing in user.User)
	if credentials.Password != "password123" {
		return "", errors.New("invalid credentials")
	}

	token := GenerateJWT(*user)
	return token, nil
}
