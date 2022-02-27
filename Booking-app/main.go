package main

import (
	"fmt"
	"strings"
)

const conferenceTicket int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings []string = []string{} //Slice

func main() { // Function that Go calls to start the whole execution

	firstName, lastName, email, userTicket := getUserInput()

	greetusers()

	for remainingTickets > 0 && len(bookings) < 50 {

		isValidName, isValidEmail, isValidTicketsNumber := validateUserInput(firstName, lastName, email, userTicket)

		// isValidCity := city == "Singpore" || city == "London" //Just checking how the OR conditional works

		if isValidName && isValidEmail && isValidTicketsNumber {

			firstNames := getFirstNames()

			fmt.Printf("The first name of bookings are: %v \n", firstNames)

			var noTicketsRemaining bool = remainingTickets == 0
			//noTicketsRemaining := remainingTickets == 0 //Alternative way

			if noTicketsRemaining {
				//End the application
				fmt.Println("Our conference is booked out. Come back next year. Thank you!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("First address you entered does not contain @ sign")
			}

			if !isValidTicketsNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}

	}

}

func greetusers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTicket, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInput(
	firstName string,
	lastName string,
	email string,
	userTicket uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2

	isValidEmail := strings.Contains(email, "@")

	isValidTicketsNumber := userTicket > 0 && userTicket <= remainingTickets

	return isValidName, isValidEmail, isValidTicketsNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint
	// Ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTicket(
	userTicket uint,
	firstName string,
	lastName string,
	email string) {
	remainingTickets = remainingTickets - userTicket
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("The whole slice: %v \n", bookings)
	fmt.Printf("The first value: %v \n", bookings[0])
	fmt.Printf("Array type: %T \n", bookings)
	fmt.Printf("Array length: %v \n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v .. \n", firstName, lastName, userTicket, email)

	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}
