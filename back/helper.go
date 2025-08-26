package main

import (
	"fmt"
	"net/http"
	"strings"
)

func greetUsers(writer http.ResponseWriter, request *http.Request) {
	// Formatted print, doesn't include newline by default, %v is the default format for variables
	fmt.Fprintf(writer, "Welcome to the %v booking application!\n", confName)
	fmt.Fprintf(writer, "We have a total of %v tickets, and %v are still available.\n", confTickets, remainingTickets)
	fmt.Fprintln(writer, "Get your tickets here to attend.")
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

func validateUserInputs(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@") && strings.Contains(email, ".")
	var isValidTicketNumber bool = userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func bookTickets(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets -= userTickets

	var userData = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("Updated list of bookings is %v\n", bookings)

	fmt.Printf("Thank you, %v %v, for booking %v ticket(s).\nYou will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are now %v tickets remaining for %v.\n", remainingTickets, confName)
}

func getFirstNames() []string {
	var firstNames []string

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}
