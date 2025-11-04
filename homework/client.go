package main

import (
	"fmt"
	"net"
)

func client(){
	var conn net.Conn
	var err error
	for conn == nil {
		conn, err = net.Dial("tcp", "localhost:10000")
	}

	if err != nil {
		fmt.Println("Eroare la conectarea la server:", err)
		return
	}
	conn.Write([]byte("Salut server!\n"))

	defer conn.Close()
}
