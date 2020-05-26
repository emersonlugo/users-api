package users

import "strconv"

type Service interface {
	GetAll() ([]*User, error)
	GetByID(userID string) (*User, error)
	Save(user *User) *User
	Update(userID string, user *User) error
	Delete(userID string) error
}

type userService struct {
	userRepository Repository
}

func NewUserService(userRepository Repository) Service {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) GetAll() ([]*User, error) {
	return s.userRepository.GetAll()
}

func (s *userService) GetByID(userID string) (*User, error) {
	return s.userRepository.GetByID(userID)
}

func (s *userService) Save(user *User) *User {
	users, _ :=  s.GetAll()
	user.ID = strconv.Itoa(len(users) + 1)
	return s.userRepository.Save(user)
}

func (s *userService) Update(userID string, user *User) error {
	return s.userRepository.Update(userID, user)
}

func (s *userService) Delete(userID string) error {
	return s.userRepository.Delete(userID)
}


