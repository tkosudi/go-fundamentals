package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")

	fmt.Println("Get your tickets he to attend")

}
