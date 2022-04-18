package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings = []string{}
	
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
	
	for {
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
		
		isValidName := len(firstName) >= 2 && len(lastName)>= 2
		isValidEmail := strings.Contains(email,"@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
		
		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			
			bookings = append(bookings, firstName + " " + lastName)
			
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
			
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
		fmt.Printf("The first names of our bookings are: %v\n", firstNames)	
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
}
