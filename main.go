package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	// lze zapsat i takto: conferenceName := "Go Conference" Takto nelze deklarovat konstanty, nelze také specificky deklarovat datový typ
	const conferenceTickets int = 50
	var remainingTickets uint = 50 // uint nemůže být záporné číslo
	var bookings []string

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName) //pointer

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your first email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaing for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// ukončit program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			// fmt.Printf("We only have %v tickets remaining, so you cant book %v tickets\n", remainingTickets, userTickets)
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address you enterd doesnt contain @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets is invalid.")
			}
		}
	}

	city := "London"

	switch city {
	case "New York":
		//execute code for booking New York conference tickets
	case "Singapore":
		//execute code for booking Singapore conference tickets
	case "London", "Berlin":
		//execute code for booking London & Berlin conference tickets
	case "Mexico City":
		//execute code for booking Mex conference tickets
	case "Hong Kong":
		//execute code for booking Hong conference tickets
	default:
		fmt.Print("No valid city selected")
	}

}
