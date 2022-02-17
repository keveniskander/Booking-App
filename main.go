package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"

	const conferenceTickets int = 50

	var remainingTickets uint = 50

	bookings := []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketCount := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// userName = "Tom"
		// userTickets = 2

		if isValidEmail && isValidName && isValidTicketCount {

			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)

			firstNames := printFirstNames(bookings)
			printFirstNames(bookings)
			fmt.Printf("The first names of bookings are %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is fully booked. Come back next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("You entered an invalid name")
			}
			if !isValidEmail {
				fmt.Println("You entered an invalid email address. Email address must contain an '@' symbol")
			}
			if !isValidTicketCount {
				fmt.Println("The number of tickets entered is invalid")
			}
			fmt.Printf("We only have %v tickets remaining. Your booking cannot be processed\n", remainingTickets)

		}

	}

}

func greetUsers(confName string, confTickets int, remTickets uint) {

	fmt.Printf("Welcome to the %v booking system\n", confName)
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", confTickets, remTickets)
	fmt.Println("Get your tickets here to attend")

}

func printFirstNames(bookings []string) []string {

	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		// fmt.Printf("Booking %v\n", booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// fmt.Scan(userName)
	// fmt.Println(remainingTickets)
	// fmt.Println(&remainingTickets)
	fmt.Println("Enter first your name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter last your name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Tickets remaining: %v\n", remainingTickets)
}
