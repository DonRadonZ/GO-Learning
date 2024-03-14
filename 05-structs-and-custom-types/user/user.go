package user

import (
	"fmt"
	"time"
	"errors"
)

type User struct {
	FirstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User
}

func (U *User) OutputUserDetails() {
	// ...

	fmt.Println(U.FirstName, U.lastName, U.birthDate)
}

func NewAdmin(email, password string) Admin {
  return Admin {
	email: email,
	password: password,
	User: User{
		FirstName: "ADMIN",
		lastName:  "ADMIN",
		birthDate: "---",
		createdAt: time.Now(),
	},
  }
}

func (U *User) ClearUserName() {
	U.FirstName = ""
	U.lastName = ""
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, last name and birthdate are required")
	}

	return &User{
		FirstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil

}
