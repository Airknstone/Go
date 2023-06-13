package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, confernce name is %T\n.", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v remaining.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// Ask use for their first name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		// Ask use for their last name
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		// Ask use for their email
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
		// Ask use for how many tickets
		fmt.Println("Enter number of tickets:  ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you cant book %v tickets.\n", remainingTickets, userTickets)
			continue
		}
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of are all our bookings: %v.\n", firstNames)

		if remainingTickets == 0 {
			//End program
			fmt.Printf("Our Conference is booked out. Come back next Year.")
			break
		}
	}
}
