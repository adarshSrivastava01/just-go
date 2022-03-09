package main

import (
	"fmt"
	"strings"
)

func main() {
	// var conferenceName string = "Go Conference"
	conferenceName := "Go Conference" // can't be done in case of constants
	const conferenceTickets uint = 50
	var remainingTickets  uint = 50
	
	fmt.Println("Welcome to", conferenceName, "booking application", conferenceTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	// var bookings = [50]string{} // length of array -> 50, type of elements -> string, {} -> Initialised Values
	// or could be replaced by
	
	var bookings []string
	// to create slice i.e. list we just remove the size from declaration -> var bookings []string
	// to add element use sliceName = append(sliceName, value)

	// for { // Statements // } -> Infinite loop

	for {
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
	
		if userTickets > remainingTickets {
			fmt.Println("Not much tickets")
			continue
		}

		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName + " " + lastName)
	
		fmt.Printf("The Whole slice: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("Slice Type: %T\n", bookings)
		fmt.Printf("Slice Length: %v\n", len(bookings))
	
		// fmt.Printf("User %v has %v Tickets", firstName, userTickets)
	
		// fmt.Printf("Type of variable can be achieved using %T\n", conferenceTickets)

		firstNames := []string{}
		// forEach loop in Go
		for _, booking := range bookings { // _ for unused variables // range -> same as enumerate in python can be used on arrays, slices, maps and more
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The First names of bookings are: %v\n", firstNames)

		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			// end the program
			fmt.Println("Our conference is booked out")
			break
		}
	}

	// Syntax for if else
	// if condtion {

	// } else if condition {

	// } else {

	// }
}