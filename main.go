package main

import "fmt"

func main() {
	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstaName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your First name")
	fmt.Scan(&firstaName)

	fmt.Println("Enter your Last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets  

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstaName, lastName, userTickets, email)

	fmt.Printf("%v ticket are still available for %v. \n", remainingTickets, conferenceName)

}