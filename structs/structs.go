package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// receiver
func (u *user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
	u.birthdate = ""
	fmt.Println("user cleared!")
}

func newUser(firstName, lastName, birthdate string) (*user, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name and birthdate are required.")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user

	appUser, err := newUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	// ... do something awesome with that gathered data!

	appUser.outputUserDetails()

	appUser.clearUserName()

	appUser.outputUserDetails()
}

func outputUserDetails(u *user) {

	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
