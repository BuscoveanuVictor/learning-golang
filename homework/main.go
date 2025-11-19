package main

import (
	"fmt"
	"os"
)

func main() {
	var choice string
	fmt.Scanln(&choice)

	switch choice {
	case "1":
		fmt.Println("Pornire server...")
		if err := runServer("127.0.0.1:10001"); err != nil {
			fmt.Printf("Eroare server: %v\n", err)
		}
	default:
		fmt.Printf("Client %d Conectat\n", os.Getpid())
		if err := runClient("127.0.0.1:10001"); err != nil {
			fmt.Printf("Eroare client: %v\n", err)
		}
	}

}

