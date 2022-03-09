package main

import "fmt"

func main() {
	// var conferenceName string = "Go Conference"
	conferenceName := "Go Conference" // can't be done in case of constants
	const conferenceTickets uint = 50
	
	fmt.Println("Welcome to", conferenceName, "booking application", conferenceTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	var userName string
	var userTickets int

	fmt.Scan(&userName) // fmt.Println(&name) -> returns the memory address

	fmt.Printf("User %v has %v Tickets", userName, userTickets)

	fmt.Printf("Type of variable can be achieved using %T\n", conferenceTickets)
}