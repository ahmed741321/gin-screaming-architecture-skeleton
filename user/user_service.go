package user

type UserService struct {
	userRepo *UserRepository
}

func NewUserService(userRepo *UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// GetUserByEmail fetches user details by email
func (s *UserService) GetUserByEmail(email string) (*User, error) {
	return s.userRepo.FindUserByEmail(email)
}
