package main

import (
	"fmt"
	"time"
)

type User struct {
  firstName string
  lastName string
  birthDate string
  createdAt time.Time	
}

func (U User) outputUserDetails(){
	// ...
	
	fmt.Println(U.firstName, U.lastName, U.birthDate)
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser User

	appUser = User{
	  firstName: userFirstName,
	  lastName: userLastName,
	  birthDate: userBirthdate,
	  createdAt: time.Now(),
	}

	// ... do something awesome with that gathered data!

	appUser.outputUserDetails()
}



func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}