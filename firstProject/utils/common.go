package utils

import (
	"fmt"
	"strings"
)

type UserData struct {
	userName        string
	email           string
	numberOfTickets uint
}

func ValidateUserInput(firstName string, lastName string, email string,
	userTickets uint, remainingTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketAmount := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketAmount
}
func GetUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var userTickets uint
	var email string

	println("Enter your first name")
	fmt.Scan(&firstName)

	println("Enter your last name")
	fmt.Scan(&lastName)

	println("How many tickets would you like to buy?")
	fmt.Scan(&userTickets)

	println("Enter email where you want tickets to be sent")
	fmt.Scan(&email)

	return firstName, lastName, email, userTickets
}

func BookTicket(remainingTickets uint, userTickets uint, userName string, email string) []UserData {

	var bookings = make([]UserData, 0)

	var userData = UserData{
		userName:        userName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	println(userName, "bought", userTickets, "tickets that will be sent to email:", email)

	remainingTickets -= userTickets

	println(remainingTickets, "are available for booking")
	return bookings
}

func GreetUser(tickets int, remainingTickets uint) {
	var appName = "Booking app"
	println("Welcome to the", appName, "buy your tikets here")
	println("Total of", tickets, "and only", remainingTickets, "are available")
}

func GetFirstNames(bookings []UserData) []string {

	fristNames := []string{}
	for _, booking := range bookings {
		fristNames = append(fristNames, booking.userName)
	}
	return fristNames
}
