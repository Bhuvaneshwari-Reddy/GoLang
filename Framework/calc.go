package framework

import "errors"

func Add(a, b int) int {
	return a + b
}

type User struct {
	Name  string
	Email string
}

func CreateUser(name, email string) (*User, error) {
	if name == "" || email == "" {
		return nil, errors.New("invalid input")
	}

	return &User{
		Name:  name,
		Email: email,
	}, nil
}
