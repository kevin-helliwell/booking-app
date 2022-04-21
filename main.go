package main

import (
	"fmt"
	"strings"
)

// Sets up variables (can't use := syntax outside of main)
const conferenceTickets int = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {
	
	
	// Greets user, gives them necessary info
	greetUsers()
	
	// Starts loop for getting tickets
	// Ceases loop if all 50 tickets are booked
	// Validates user input and number of tickets booked by user
	for {
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)
		
	if isValidName && isValidEmail && isValidTicketNumber {
		
		bookTicket(userTickets, firstName, lastName, email)
			
		// Get first names and print them
		firstNames := getFirstNames()
		fmt.Printf("The first names of our bookings are: %v\n",firstNames)	
	}
	if remainingTickets == 0 {
		// end program
		fmt.Println("Our conference is sold out. Come back next year.")
		break
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered does not contain @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid.")
			}
	}	
}
	// Switch statement example for handling/consolidating control flow through cases
	// city := "London"
	
	// switch city {
	// 	case "New York":
	// 		// execute code for booking New York conference tickets
	// 	case "Singapore", "Hong Kong":
	// 		// execute code for booking Singapore & Hong Kong conference tickets
	// 	case "London", "Berlin":
	// 		// execute code for booking London & Berlin conference tickets
	// 	case "Mexico City":
	// 		// execute code for booking Mexico City conference tickets
	// 	default:
	// 		fmt.Print("No valid city selected.")
	// }

}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
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

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName)>= 2
	isValidEmail := strings.Contains(email,"@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		
		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
		
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)
		
		return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)
			
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName,lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n",remainingTickets, conferenceName)
}