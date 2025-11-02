package main

import(
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net"
)

type FruitQuantities struct {
	Mere    float64 `json:"mere"`
	Pere    float64 `json:"pere"`
	Piersici float64 `json:"piersici"`
	Capsuni  float64 `json:"capsuni"`
}


func ex1(){
	file, err := os.Open("cantitati_fructe.json")
	if err != nil {
		fmt.Println("Eroare la deschiderea fisierului:", err)
		return
	}

	dataBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Eroare la citirea fisierului:", err)
		return
	}

	fmt.Println("Fisierul a fost citit cu succes.")
	fmt.Println("Numele fisierului:", file.Name())
	fmt.Println("Continutul fisierului:", string(dataBytes))

	// Avem datele in dataBytes 
	// continutul fisierului are strucutura JSON
	// putem sa le procesam mai departe


	// Metoda 1: folosim Unmarshal pentru a converti datele JSON intr-o structura Go
	var fruitQuantities FruitQuantities
	err = json.Unmarshal(dataBytes, &fruitQuantities)
	if err != nil {
		fmt.Println("Eroare la deserializare JSON:", err)
		return
	}

	fmt.Println("Datele deserializate din JSON:")
	fmt.Println("Mere:", fruitQuantities.Mere)
	fmt.Println("Pere:", fruitQuantities.Pere)
	fmt.Println("Piersici:", fruitQuantities.Piersici)
	fmt.Println("Capsuni:", fruitQuantities.Capsuni)


	// Metoda 2: folosim un map pentru a stoca datele JSON
	var fruitMap map[string]float64
	err = json.Unmarshal(dataBytes, &fruitMap)
	if err != nil {
		fmt.Println("Eroare la deserializare JSON:", err)
		return
	}

	fmt.Println("Datele deserializate din JSON (metoda 2):")
	for fruit, quantity := range fruitMap {
		fmt.Printf("%s: %f\n", fruit, quantity)
	}

	defer file.Close()
}

func ex2(){
	listener, err := net.Listen("tcp", ":9090")

	if err != nil {
		fmt.Println("Eroare la pornirea serverului:", err)
		return
	}

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Eroare la acceptarea conexiunii:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Conexiune acceptata de la:", conn.RemoteAddr())


	message := "Bine ati venit la server!\n"
	_, err = conn.Write([]byte(message))

	b := make([]byte, 1024)
	_, err = conn.Read(b)
	if err != nil {
		fmt.Println("Eroare la citirea datelor:", err)
		return
	}
	fmt.Println("Mesaj de la client:", string(b))

	defer listener.Close()
}

func ex3(){
	conn, err := net.Dial("tcp", ":9090")
	if err != nil {
		fmt.Println("Eroare la conectarea la server:", err)
		return
	}


	data := make([]byte, 1024)
	n, err := conn.Read(data)
	if err != nil {
		fmt.Println("Eroare la citirea datelor:", err)
		return
	}

	fmt.Println("Mesaj de la server:", string(data[:n]))

	conn.Write([]byte("Multumesc server!\n"))

	defer conn.Close()
}