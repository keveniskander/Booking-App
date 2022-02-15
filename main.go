package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"

	const conferenceTickets int = 50

	var remainingTickets uint = 50

	bookings := []string{}

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketCount := userTickets > 0 && userTickets < remainingTickets
		// userName = "Tom"
		// userTickets = 2

		if isValidEmail && isValidName && isValidTicketCount {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			// fmt.Printf("The whole Slice: %v\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("Slice type: %T\n", bookings)
			// fmt.Printf("Slice length: %v\n", len(bookings))

			// fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("Tickets remaining: %v\n", remainingTickets)
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				// fmt.Printf("Booking %v\n", booking)
				firstNames = append(firstNames, names[0])
			}
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
