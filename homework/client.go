package main

import (
	"fmt"
	"net/rpc"
	"os"
	"log"
	"strings"
)

var reply Reply
var args Args

func runClient(addr string) error {
	client, err := rpc.Dial("tcp", addr)
	if err != nil {
		return fmt.Errorf("nu pot face dial catre server: %w", err)
	}
	defer client.Close()

	data, err := os.ReadFile("input4.txt")
	if err != nil {
		log.Fatal(err)
	}
	

	dataStr := string(data)
	dataSplit := strings.Split(strings.TrimSpace(dataStr), " ")

	args := Args{
		Data: dataSplit[1:],
	}
	
	switch dataSplit[0] {
		case "1":
			fmt.Println("Clientul a selectat cerinta 1")
			if err := client.Call("Procedures.Words", &args, &reply); err != nil {
				return fmt.Errorf("apel RPC esuat: %w", err)
			}
	
		case "3":
			fmt.Println("Clientul a selectat cerinta 3")
			if err := client.Call("Procedures.Numbers", &args, &reply); err != nil {
				return fmt.Errorf("apel RPC esuat: %w", err)
			}

		case "4":
			fmt.Println("Clientul a selectat cerinta 4")
			if err := client.Call("Procedures.Range", &args, &reply); err != nil {
				return fmt.Errorf("apel RPC esuat: %w", err)
			}
	}


	fmt.Println("Rezultat:", reply.Result)

	// fmt.Println("Rezultat:", reply.Result)
	return nil
}

