package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	// lze zapsat i takto: conferenceName := "Go Conference" Takto nelze deklarovat konstanty, nelze také specificky deklarovat datový typ
	const conferenceTickets int = 50
	var remainingTickets uint = 50 // uint nemůže být záporné číslo

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) //pointer

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your first email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

}
