package users

import (
	"errors"
	"log"
)

const msgUserNotFound = "user %s not found"

var ErrorUserNotFound = errors.New("user not found")
var ErrorUsersNotFound = errors.New("users not found")

type Repository interface {
	GetAll() ([]*User, error)
	GetByID(userID string) (*User, error)
	Save(user *User) *User
	Update(userID string, user *User) error
	Delete(userID string) error
}

type memoryRepository struct {
	users []*User
}

func NewInMemoryUserRepository(users ...*User) Repository {
	return &memoryRepository{
		users: append([]*User{}, users...),
	}
}

func (r *memoryRepository) GetAll() ([]*User, error) {
	if len(r.users) == 0 {
		log.Print("users not found")
		return r.users, ErrorUsersNotFound
	}
	return r.users, nil
}

func (r *memoryRepository) GetByID(userID string) (*User, error) {
	for _, user := range r.users {
		if user.ID == userID {
			return user, nil
		}
	}
	log.Printf(msgUserNotFound, userID)
	return nil, ErrorUserNotFound
}

func (r *memoryRepository) Save(user *User) *User {
	r.users = append(r.users, user)
	return user
}

func (r *memoryRepository) Update(userID string, userToUpdate *User) error {
	for index, user := range r.users {
		if user.ID == userID {
			r.users[index] = userToUpdate
			return nil
		}
	}
	log.Printf(msgUserNotFound, userID)
	return ErrorUserNotFound
}

func (r *memoryRepository) Delete(userID string) error {
	if len(r.users) == 0 {
		log.Printf(msgUserNotFound, userID)
		return ErrorUserNotFound
	}

	users := r.users[:0]
	for _, user := range r.users {
		if user.ID != userID {
			users = append(users, user)
		}
	}

	if len(r.users) == len(users) {
		log.Printf(msgUserNotFound, userID)
		return ErrorUserNotFound
	}

	r.users = users
	return nil
}
