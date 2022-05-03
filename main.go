package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {

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

		isValidName := len(firstaName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets -= userTickets
			bookings = append(bookings, firstaName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v. \n", firstaName, lastName, userTickets, email)

			fmt.Printf("%v ticket are still available for %v. \n", remainingTickets, conferenceName)

			firstaNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstaNames = append(firstaNames, names[0])
			}

			fmt.Printf("The first names of bookings are %v \n", firstaNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("The email you entered doesn't have an @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is invalid")
			}
		}
	}

}
