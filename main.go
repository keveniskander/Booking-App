package main

import "fmt"

func main() {

	conferenceName := "Go Conference"

	const conferenceTickets = 50

	var remainingTickets = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int

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
	// userName = "Tom"
	// userTickets = 2

	// fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

}
