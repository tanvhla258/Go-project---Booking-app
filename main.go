package main

import "fmt"

func main() {

	var confernenceName = "Go Conference"
	const confernenceTickets = 50
	var remainingTickets = 50
	fmt.Println("Welcome to", confernenceName, "booking application")
	fmt.Println("We have total of", confernenceTickets, "tickets and", remainingTickets, "available")

	fmt.Println("Get your tickets here to attend")

	fmt.Println(confernenceName)
}
