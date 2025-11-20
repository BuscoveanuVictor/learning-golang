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
	log.Println("Server a primit requestul.")
	log.Println("Server proceseaza datele.")

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
	log.Printf("Server trimite %v catre client.\n", reply.Result)

    return  nil
}


func (Procedures) Numbers(args *Args, reply *Reply) error {
	log.Println("Server a primit requestul.")
	log.Println("Server proceseaza datele.")
	
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
	log.Printf("Server trimite %v catre client.\n", reply.Result)

    return nil
}

func (Procedures) Range(args *Args, reply *Reply) error {
	log.Println("Server a primit requestul.")
	log.Println("Server proceseaza datele.")
	
	var data []int
	
	for _, d := range(args.Data){
		num, _ := strconv.Atoi(d)		
		data = append(data, num)
	}


	a := data[0]
	b := data[1]

	log.Println("Interval: ", a, " ", b)

	resSum := 0
	n := 0	
	for _, val := range data[3:] {

		sum := 0
		num := val
		for num > 0 {
			digit := num % 10
			sum += digit
			num /= 10
		}
		log.Println("Suma cifrelor: ", sum)
		if sum >= a && sum <= b {
			resSum += val
			n += 1
		}
	}
	log.Println("Suma numerelor care indeplinesc conditia: ", resSum)

	if n == 0 {
		reply.Result = []string{"0"}
		log.Printf("Server trimite %v catre client.\n", reply.Result)
		return nil
	}

	reply.Result = []string{strconv.Itoa(resSum / n)}
	log.Printf("Server trimite %v catre client.\n", reply.Result)
	return nil

}

func (Procedures) PrimeDigits(args *Args, reply *Reply) error {
	log.Println("Server a primit requestul.")
	log.Println("Server proceseaza datele.")
	totalDigits := 0

	for _, d := range args.Data {
		num, _ := strconv.Atoi(d)
		
		if isPrime(num) {
			totalDigits += digitCount(num)
		}
	}

	reply.Result = []string{strconv.Itoa(totalDigits)}
	log.Printf("Server trimite %v catre client.\n", reply.Result)
	return nil
}

func (Procedures) DoubleFirstDigitSum(args *Args, reply *Reply) error {
	log.Println("Server a primit requestul.")
	log.Println("Server proceseaza datele.")
	sum := 0

	for _, d := range args.Data {
		num, _:= strconv.Atoi(d)

		sum += doubleFirstDigitValue(num)
	}

	reply.Result = []string{strconv.Itoa(sum)}
	log.Printf("Server trimite %v catre client.\n", reply.Result)
	return nil
}

func digitSum(n int) int {
	if n < 0 {
		n = -n
	}
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func digitCount(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
	}
	count := 0
	for n > 0 {
		count++
		n /= 10
	}
	return count
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func doubleFirstDigitValue(n int) int {
	if n == 0 {
		return 0
	}

	first := firstDigit(n)
	digits := digitCount(n)

	restPow := 1
	for i := 0; i < digits-1; i++ {
		restPow *= 10
	}
	rest := n % restPow

	return first*(restPow*10) + first*restPow + rest
}

func firstDigit(n int) int {
	if n < 0 {
		n = -n
	}
	for n >= 10 {
		n /= 10
	}
	return n
}