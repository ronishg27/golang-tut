package users

import (
	"strings"
)

type User struct {
	Name    string
	Age     int
	Email   string
	Address string
	Status  bool
}

// users := []User{}

func (u User) CapitalizedUserName() string {
	return strings.ToUpper(u.Name)
}

func NewUser(name string, age int, email string, address string, status bool) User {
	return User{
		Name:    name,
		Age:     age,
		Email:   email,
		Address: address,
		Status:  status,
	}
}
