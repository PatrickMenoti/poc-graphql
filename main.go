package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/PatrickMenoti/poc-graphql/options"
)

var tokenvalue string

func main() {

	fmt.Println("Select an option:")
	fmt.Println("1. Top status codes")
	fmt.Println("2. Top user agent")
	fmt.Println("3. Total requests")

	var choice int
	fmt.Print("Enter your choice: ")
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	flag.StringVar(&tokenvalue, "token", "", "API token for authorization")
	// Parse the command-line flags
	flag.Parse()

	// Check if the token is provided
	if tokenvalue == "" {
		fmt.Println("Please provide a valid API token using the -token flag.")
		return
	}

	switch choice {
	case 1:
		options.Option1(tokenvalue)
	case 2:
		options.Option2(tokenvalue)
	case 3:
		options.Option3(tokenvalue)
	default:
		fmt.Println("Invalid choice. Please enter 1 or 2.")
	}

}
