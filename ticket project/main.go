package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets,and %v tickets are ramaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Println("touch is working")

	for {

		var bookings []string
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your Last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your Email Id: ")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets to be booked: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We have only %v tickets ramining \n", remainingTickets)
			continue
		}
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The whole slice : %v\n", bookings)
		fmt.Printf("The first value : %v\n", bookings[0])
		fmt.Printf("The slice type  : %T\n", bookings)
		fmt.Printf("slice Length  : %v\n", len(bookings))

		fmt.Printf("Customer with name %v , have booked %v tickets for the conference.\n", firstName, userTickets)

		fmt.Printf("Thank you for booking the ticket through %v\n , you will receive tickets in your %v mail id.\n", conferenceName, email)
		fmt.Printf("%v tickets ramining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}

		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The firstnames of bookings are : %v\n", firstNames)

		if remainingTickets == 0 {
			//end the programme
			fmt.Println("Our conference tickets are sold out , please come back next year")
			break
		}
	}

}
