package main

import (
	"fmt"
	"os"
)

func main() {
	var choice string
	
	var req string

	for{
		fmt.Println("Rulare server(1) sau client(2)? Iesire(0)")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			fmt.Println("Pornire server...")
			if err := runServer("127.0.0.1:12000"); err != nil {
				fmt.Printf("Eroare server: %v\n", err)
			}
		case "2":
			fmt.Printf("Client %d Conectat\n", os.Getpid())

			fmt.Println("Ce cerinta doriti sa rulati? (1, 3, 4, 8 sau 12)")
			fmt.Scanln(&req)

			if(req != "1" && req != "3" && req != "4" && req != "8" && req != "12"){
				fmt.Println("Cerinta invalida!")
				continue
			}

			if err := runClient("127.0.0.1:12000", req, os.Getpid()); err != nil {
				fmt.Printf("Eroare client: %v\n", err)
			}

		case "0":
			fmt.Println("Iesire...")
			return
		}
	}
}
