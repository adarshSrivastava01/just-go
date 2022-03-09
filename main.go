package main

import "fmt"

func main() {
	// var conferenceName string = "Go Conference"
	conferenceName := "Go Conference" // can't be done in case of constants
	const conferenceTickets uint = 50
	var remainingTickets  uint = 50
	
	fmt.Println("Welcome to", conferenceName, "booking application", conferenceTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	// var bookings = [50]string{} // length of array -> 50, type of elements -> string, {} -> Initialised Values
	// or could be replaced by
	
	var bookings [50]string

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter the first Name: ")
	fmt.Scan(&firstName) // fmt.Println(&name) -> returns the memory address
	fmt.Println("Enter the last Name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter the Email: ")
	fmt.Scan(&email)
	fmt.Println("Enter the No of Tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings[0] = firstName + " " + lastName

	fmt.Printf("The Whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array Type: %T\n", bookings)
	fmt.Printf("Array Length: %v\n", len(bookings))

	// fmt.Printf("User %v has %v Tickets", firstName, userTickets)

	// fmt.Printf("Type of variable can be achieved using %T\n", conferenceTickets)
}