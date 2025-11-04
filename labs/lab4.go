package main

import (
	"encoding/json"
	"fmt"
	"encoding/xml"
	"os"
	"net/url"
	"time"
)

type Person struct {
	Nume  string `json:"nume"`
	Grupa string `json:"grupa"`	
}


func lab4_ex1(){
	p := Person{
		Nume:  "Popescu",
		Grupa: "333",
	}

	fmt.Println("Datele initiale:", p.Nume, p.Grupa)


	// Atentie !
	// Tu chiar daca ai folosit Marshal
	// datele returnate sunt in format de bytes
	// adica nu ti converteste intr-o stuctura de tip json !
	dateJson, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Eroare la serializare:", err)
		return
	}

	fmt.Println("Datele in format JSON:")
	fmt.Println(string(dateJson))

}


type Coffee struct {
    XMLName xml.Name `xml:"coffee"`
    Name    string   `xml:"name"`
    Origin  string   `xml:"origin"`
}

func lab4_ex2(){
	coffee := Coffee{
		Name:   "Arabica",
		Origin: "Ethiopia",
	}

	fmt.Println("Datele initiale:", coffee.XMLName, coffee.Origin)

	xmlData, err := xml.MarshalIndent(coffee, "", "  ")
	if err != nil {
		fmt.Println("Eroare la serializare XML:", err)
		return
	}

	fmt.Println("Datele in format XML:")
	fmt.Println(string(xmlData))


	os.WriteFile("coffee.xml", xmlData, 0644)



	// Deserializare XML
	var coffee2 Coffee
	err = xml.Unmarshal(xmlData, &coffee2)
	if err != nil {
		fmt.Println("Eroare la deserializare XML:", err)
		return
	}

	fmt.Println("Datele deserializate din XML:")
	fmt.Println("Nume:", coffee2.Name)
	fmt.Println("Origine:", coffee2.Origin)

}

func lab4_ex3(){
	parsedUrl, err := url.Parse("https://example.com:8080/path?item=1#fragment")
	if err != nil {
		fmt.Println("Eroare la parsare URL:", err)
		return
	}

	fmt.Println("Protocol:", parsedUrl.Scheme)
	fmt.Println("Host:", parsedUrl.Host)
	fmt.Println("Port:", parsedUrl.Port())
	fmt.Println("Path:", parsedUrl.Path)
	fmt.Println("Query:", parsedUrl.Query().Get("item"))
	fmt.Println("Fragment:", parsedUrl.Fragment)
}

func ex4(){	
	currentTime := time.Now()
	fmt.Println("Data si ora curenta (format implicit):", currentTime.String())
	fmt.Println("Data si ora curenta:", currentTime.Format("15:04:05 02-Jan-2006"))


	timeStr := "21:30:00 15-Aug-2023"
	parsedTime, err := time.Parse("15:04:05 02-Jan-2006", timeStr)
	if err != nil {
		fmt.Println("Eroare la parsare data si ora:", err)
		return
	}
	fmt.Println("Data si ora parsata:", parsedTime.Add(10*time.Minute).Format("15:04:05 02-Jan-2006"))


	fmt.Println("Timestamp Unix:", currentTime.Unix())
}
