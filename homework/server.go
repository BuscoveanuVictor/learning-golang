package main

import (
    "fmt"
    "log"
    "net"
    "net/rpc"
    "strconv"
)

type Args struct {
	Data    []string
}

type Reply struct {
	Result  []string
}

type Procedures struct{}

func runServer(addr string) error {
	
    if err := rpc.RegisterName("Procedures", &Procedures{}); err != nil {
        return fmt.Errorf("nu pot inregistra serviciul: %w", err)
    }

    listener, err := net.Listen("tcp", addr)
    if err != nil {
        return fmt.Errorf("nu pot asculta pe %s: %w", addr, err)
    }
	log.Printf("Server asculta pe %s", addr)

	defer listener.Close()

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("eroare acceptare conexiune: %v", err)
            continue
        }
        go rpc.ServeConn(conn)
    }

	return nil
	
}

func (Procedures) Words(args *Args, reply *Reply) error {
	words := args.Data

	firstLen := len(words[0])

    for _, w := range words {
        if len(w) != firstLen {
            return fmt.Errorf("cuvintele nu au lungimi egale!")
        }
    }

    res := make([]string, firstLen)

    for i := 0; i < firstLen; i++ {
        for _, w := range words {
            res[i] += string(w[i])
        }
    }

	reply.Result = res

    return  nil
}


func (Procedures) Numbers(args *Args, reply *Reply) error {
	
	sum:= 0

	for _, str := range args.Data {
		num, _ := strconv.Atoi(str)
		
		x := 0
        for num > 0 {
            digit := num % 10
            x = x * 10 + digit
            num /= 10
        }
        sum += x
	}

	// Converteste int la string
	reply.Result = []string{(strconv.Itoa(sum))}

    return nil
}

func (Procedures) Range (args *Args, reply *Reply) error {
	
	var data []int
	
	for _, d := range(args.Data){
		num, _ := strconv.Atoi(d)
		
		data = append(data, num)
	}

	a := data[1]
	b := data[2]

	log.Println("Interval: ", a, " ", b)

	resSum := 0
	n := 0
	for _, num := range(data[4:]){
		sum:=0
		for num>0{
			digit := num % 10
			sum	+= digit
			num /= 10
		}
		log.Println("Suma cifrelor: ", sum)
		if sum >= a && sum <= b {
			resSum += num
			n += 1
		}
	}
	log.Println("Suma numerelor care indeplinesc conditia: ", resSum)

	reply.Result = []string{strconv.Itoa(resSum/n)}
	return nil

}