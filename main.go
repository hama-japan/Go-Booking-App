package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	
	var firstName string
	var userTickets int
	// ask user for their name
	fmt.Println("Enter your fiest name: ")
	fmt.Scan(&firstName)

	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)
}