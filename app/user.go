package app

import "fmt"

// Struct for a user
type User struct {
	Name     string
	Login    string
	Password string
}

// Struct to keep app users
type Users struct {
	Users []User
}

func (s *Users) FindUser(login string) (User, error) {
	for _, u := range s.Users {
		if u.Login == login {
			return u, nil
		}
	}
	return User{}, fmt.Errorf("no users found")
}
