package auth

import (
	"fmt"
	handler "p1/internal/handlers"
	"p1/internal/router"
	"p1/internal/service"
)

// Define a custom type for login enums
type LoginType int

// Enumerated constants using iota
const (
	Client LoginType = iota
	User
	Unknown
)

// Convert string input to the corresponding LoginType enum
func getLoginType(login string) LoginType {
	switch login {
	case "client":
		return Client
	case "user":
		return User
	default:
		return Unknown
	}
}

func InitLogin() {
	route := router.NewRouter()
	fmt.Printf("Enter the client or user login: ")
	var login string

	fmt.Scanf("%s", &login) 
	fmt.Printf("entered here\n")
	// Get the enum for the input
	loginType := getLoginType(login)

	// Use a switch statement to handle login types
	switch loginType {
	case Client:
		fmt.Println("You have selected client login.")
		// Add your client login logic here
		clientService,_ := service.NewClientService()
		clienthandler := handler.NewClientHandler(clientService)
		router.StarctClientServer(clienthandler,route)
		router.StartServer(route)
		
	case User:
		fmt.Println("You have selected user login.")
		// Add your user login logic here
	default:
		fmt.Println("Invalid login type. Please enter either 'client' or 'user'.")
	}
}
