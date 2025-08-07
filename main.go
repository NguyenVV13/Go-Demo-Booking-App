package main

import "fmt"

// typehint is after variable name instead of before, unlike other languages
const confTickets uint = 50

var confName string = "Go Conference"
var remainingTickets uint = confTickets
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

func main() {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	greetUsers()

	for remainingTickets > 0 {
		firstName, lastName, email, userTickets = getUserInputs(firstName, lastName, email, userTickets)
		isValidName, isValidEmail, isValidTicketNumber := validateUserInputs(firstName, lastName, email, userTickets)

		if !isValidName {
			fmt.Println("Your first and last names must have at least 2 characters each")
			continue
		}
		if !isValidEmail {
			fmt.Println("Your email must contain @ and a domain (ex: gmail.com)")
			continue
		}
		if !isValidTicketNumber {
			fmt.Printf("Only %v tickets remain, so you can't book %v tickets.\n", remainingTickets, userTickets)
			continue
		}

		// else ...
		bookTickets(firstName, lastName, email, userTickets)

		fmt.Printf("These are the first names for all existing bookings: %v\n", getFirstNames())
	}

	// End program
	fmt.Println("The conference is booked out. Please come check again next year.")
}
