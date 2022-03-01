package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTicket int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50

// var bookings []string = []string{} //Slice
var bookings = make([]UserData, 0) //Initializing a list of maps

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() { // Function that Go calls to start the whole execution

	firstName, lastName, email, userTicket := getUserInput()

	greetusers()

	isValidName, isValidEmail, isValidTicketsNumber := helper.ValidateUserInput(
		firstName, lastName, email, userTicket, remainingTickets)

	// isValidCity := city == "Singpore" || city == "London" //Just checking how the OR conditional works

	if isValidName && isValidEmail && isValidTicketsNumber {

		bookTicket(userTicket, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTicket, firstName, lastName, email)

		firstNames := getFirstNames()

		fmt.Printf("The first name of bookings are: %v \n", firstNames)

		var noTicketsRemaining bool = remainingTickets == 0
		//noTicketsRemaining := remainingTickets == 0 //Alternative way

		if noTicketsRemaining {
			//End the application
			fmt.Println("Our conference is booked out. Come back next year. Thank you!")
			// break
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

	wg.Wait()
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

		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
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

	//Creating a map for a user
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["NumberOfTickets"] = strconv.FormatUint(uint64(userTicket), 10)

	//Better way to creating a map for a user
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTicket,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("The whole slice: %v \n", bookings)
	fmt.Printf("The first value: %v \n", bookings[0])
	fmt.Printf("Array type: %T \n", bookings)
	fmt.Printf("Array length: %v \n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v .. \n", firstName, lastName, userTicket, email)

	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)
}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tiekcts for %v %v", userTicket, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("###############")

	wg.Done()
}
