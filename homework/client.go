package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

var reply Reply
var args Args



func runClient(addr string, req string, clientName int) error {
	client, err := rpc.Dial("tcp", addr)
	if err != nil {
		return fmt.Errorf("nu pot face dial catre server: %w", err)
	}
	defer client.Close()

	fmt.Printf("Client %d Conectat.\n", clientName)

	data, err := os.ReadFile(fmt.Sprintf("input%s.txt", req))
	if err != nil {
		log.Fatal(err)
	}
	

	dataStr := string(data)
	dataSplit := strings.Split(strings.TrimSpace(dataStr), " ")

	args := Args{
		Data: dataSplit,
	}

	m := make(map[string]string)
	m["1"] = "Words"
	m["3"] = "Numbers"
	m["4"] = "Range"
	m["8"] = "PrimeDigits"
	m["12"] = "DoubleFirstDigitSum"

	procedureToCall, ok := m[req]
	if !ok {
		return fmt.Errorf("cerinta %s nu este suportata", req)
	}

	fmt.Printf("Client %d a facut request cu datele: %s.\n", clientName, strings.Join(args.Data, " "))

	fmt.Println("Clientul a selectat cerinta: ", req)
	if err := client.Call("Procedures."+procedureToCall, &args, &reply); err != nil {
		return fmt.Errorf("apel RPC esuat: %w", err)
	}

	fmt.Printf("Client %d a primit raspunsul: %v.\n", clientName, reply.Result)
	fmt.Println("Rezultat:", reply.Result)

	return nil
}

