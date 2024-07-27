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

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for remainingTickets > 0 && len(bookings) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name
		fmt.Println("Enter your fiest name: ")
		fmt.Scan(&firstName)
	
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
	
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
	
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName)
		
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		
			firseNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				fitstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
	
			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else if userTickets == 0 {
			fmt.Println("うりきれです")
			break
		} else {
			if !isValidName {
				fmt.Println("名前が無効です")
			}
			if !isValidEmail {
				fmt.Println("メールアドレスがむこうです")
			}
			if !isValidTicketNumber {
				fmt.Println("チケット枚数が無効です")
			}
		}

	}

}
