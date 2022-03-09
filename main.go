package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	
	fmt.Println("Welcome to", conferenceName, "booking application", conferenceTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2

	fmt.Printf("User %v has %v Tickets", userName, userTickets)

	fmt.Printf("Type of variable can be achieved using %T\n", conferenceTickets)
}