package main

import (
	"booking-app/utils"
	"fmt"
	"time"
)

func main() {
	const tickets int = 50
	var remainingTickets uint = 50
	utils.GreetUser(tickets, remainingTickets)

	for {
		firstName, lastName, email, userTickets := utils.GetUserInput()
		isValidName, isValidEmail, isValidTicketAmount := utils.ValidateUserInput(firstName, lastName,
			email, userTickets, remainingTickets)

		if !isValidName {
			println("First name is too short. Please try again")
			continue
		}

		if !isValidEmail {
			println("Invalid email. Please try again")
			continue
		}

		if !isValidTicketAmount {
			println("Invalid amount of tickets. Please try again")
			continue
		}

		var userName = firstName + " " + lastName
		bookings := utils.BookTicket(remainingTickets, userTickets, userName, email)
		go sendTicket(userTickets, userName)
		firstNames := utils.GetFirstNames(bookings)
		fmt.Printf("These are all the bookings %v \n", firstNames)

		if remainingTickets == 0 {
			println("Sold out")
			break
		}
	}

}

func sendTicket(userTickets uint, userName string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v", userTickets, userName)
	println("############")
	println("Sending the", ticket, "to the customer", userName)
	println("############")
}
