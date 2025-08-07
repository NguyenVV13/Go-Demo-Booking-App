package main

import (
	"fmt"
	"strings"
)

// typehint is after variable name instead of before, unlike other languages
const confTickets uint = 50

var confName string = "Go Conference"
var remainingTickets uint = confTickets
var bookings = []string{}

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
		bookTickets(firstName, lastName, userTickets)

		fmt.Printf("Thank you, %v %v, for booking %v ticket(s).\nYou will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
		fmt.Printf("There are now %v tickets remaining for %v.\n", remainingTickets, confName)
		fmt.Printf("These are the first names for all existing bookings: %v\n", getFirstNames())
	}

	// End program
	fmt.Println("The conference is booked out. Please come check again next year.")
}

func greetUsers() {
	// Formatted print, doesn't include newline by default, %v is the default format for variables
	fmt.Printf("Welcome to the %v booking application!\n", confName)
	fmt.Printf("We have a total of %v tickets, and %v are still available.\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getUserInputs(firstName string, lastName string, email string, userTickets uint) (string, string, string, uint) {
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter desired number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(firstName string, lastName string, userTickets uint) {
	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

}

func getFirstNames() []string {
	var firstNames []string

	for _, booking := range bookings {
		var names []string = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}
