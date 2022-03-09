package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	
	fmt.Println("Welcome to", conferenceName, "booking application", conferenceTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)

}