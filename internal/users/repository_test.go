package users

import (
	"reflect"
	"testing"
)

func Test_MemoryRepository_GetAll(t *testing.T) {
	testCases := []struct {
		name          string
		users         []*User
		expectedUsers  []*User
		expectedError error
	}{
		{
			name:          "Should return error when not found users",
			users:         []*User{},
			expectedError: ErrorUsersNotFound,
		},
		{
			name:          "Should return all users",
			users:         mockUsers(),
			expectedError: nil,
		},
	}

	for i, tc := range testCases {
		repository := NewInMemoryUserRepository(tc.users...)
		users, err := repository.GetAll()

		if !reflect.DeepEqual(users, tc.users){
			t.Errorf("%d:%s: got %v; want %v", i, tc.name, users, tc.users)
		}

		if err != tc.expectedError {
			t.Errorf("%d:%s: got %q; want %q", i, tc.name, err, tc.expectedError)
		}
	}
}

func Test_MemoryRepository_GetByID(t *testing.T) {
	testCases := []struct {
		name          string
		userID        string
		users         []*User
		expectedUser  *User
		expectedError error
	}{
		{
			name:          "Should return error when not found user by id",
			userID:        "1",
			users:         []*User{},
			expectedUser:  nil,
			expectedError: ErrorUserNotFound,
		},
		{
			name:          "Should return user for searched id",
			userID:        "1",
			users:         mockUsers(),
			expectedUser:  &User{ID: "1", Name: "Emerson Lugo"},
			expectedError: nil,
		},
	}

	for i, tc := range testCases {
		repository := NewInMemoryUserRepository(tc.users...)
		user, err := repository.GetByID(tc.userID)

		if !reflect.DeepEqual(user, tc.expectedUser) {
			t.Errorf("%d:%s: got %q; want %q", i, tc.name, user, tc.expectedUser)
		}

		if err != tc.expectedError {
			t.Errorf("%d:%s: got %q; want %q", i, tc.name, err, tc.expectedError)
		}
	}
}

func Test_MemoryRepository_Save(t *testing.T) {
	users := mockUsers()
	initialLength := len(users)
	repository := &memoryRepository{users: users}
	newUser := &User{ID: "3", Name: "Patrick"}

	if user := repository.Save(newUser); user == nil {
		t.Errorf("%s: got %q; want %q", "save user", user, newUser)
	}

	if initialLength == len(repository.users) {
		t.Errorf("%s: got %q; want %q", "save user", initialLength+1, len(repository.users))
	}
}

func Test_MemoryRepository_Update(t *testing.T) {
	testCases := []struct {
		name          string
		userID        string
		users         []*User
		userUpdate    *User
		expectedError error
	}{
		{
			name:          "Should return error when not found users to update",
			userID:        "101",
			users:         mockUsers(),
			userUpdate:    &User{ID: "101", Name: "Patrick"},
			expectedError: ErrorUserNotFound,
		},
		{
			name:          "Should return nil when update user successfully",
			userID:        "1",
			users:         mockUsers(),
			userUpdate:    &User{ID: "1", Name: "Emerson Lugo", Email: "emersonlugo@gmail.com"},
			expectedError: nil,
		},
	}

	for i, tc := range testCases {
		repository := NewInMemoryUserRepository(tc.users...)
		if err := repository.Update(tc.userID, tc.userUpdate); err != tc.expectedError {
			t.Errorf("%d:%s: got %q; want %q", i, tc.name, err, tc.expectedError)
		}
	}
}

func Test_MemoryRepository_Delete(t *testing.T) {
	testCases := []struct {
		name          string
		userID        string
		users         []*User
		expectedError error
	}{
		{
			name:          "Should return error when not found users to delete",
			userID:        "101",
			users:         []*User{},
			expectedError: ErrorUserNotFound,
		},
		{
			name:          "Should return error when not found user by id to delete",
			userID:        "3",
			users:         mockUsers(),
			expectedError: ErrorUserNotFound,
		},
		{
			name:          "Should return nil when delete user successfully",
			userID:        "1",
			users:         mockUsers(),
			expectedError: nil,
		},
	}

	for i, tc := range testCases {
		repository := NewInMemoryUserRepository(tc.users...)
		if err := repository.Delete(tc.userID); err != tc.expectedError {
			t.Errorf("%d:%s: got %q; want %q", i, tc.name, err, tc.expectedError)
		}
	}
}

func mockUsers() []*User {
	return []*User{
		{
			ID:   "1",
			Name: "Emerson Lugo",
		},
		{
			ID:   "2",
			Name: "Mister Golang",
		},
	}
}
