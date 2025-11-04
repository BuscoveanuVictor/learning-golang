package main

import (
	"fmt"
	"net"
)
func server(){
	listener, err := net.Listen("tcp", "localhost:10000")
	if err != nil {
		fmt.Println("Eroare la pornirea serverului:", err)
		return
	}
	fmt.Println("Serverul a pornit si asteapta conexiuni...")
	

	for{
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Eroare la acceptarea conexiunii:", err)
			continue
		}
		buffer := make([]byte, 1024)
		conn.Read(buffer)
		fmt.Println("Mesaj de la client:", string(buffer))
	}

}
