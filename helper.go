package main

import (
	"fmt"
	"strings"
)

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

func validateUserInputs(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@") && strings.Contains(email, ".")
	var isValidTicketNumber bool = userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
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
