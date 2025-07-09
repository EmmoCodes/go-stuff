package user

import (
	"errors"
	"fmt"
	"time"
)

type TUser struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	TUser
}

func (user *TUser) OutputUserDetails() {
	// ...
	fmt.Println(user.firstName, user.lastName, user.birthdate)
}

func (user *TUser) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		TUser: TUser{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}

func New(firstName, lastName, birthdate string) (*TUser, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name and birthdate are required!")
	}

	return &TUser{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
