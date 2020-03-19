package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {

	interfaceToConnect := flag.String("interface", "0.0.0.0:4000", "an interface to startup")
	mode := flag.Int("mode", 1, "1 - load, 2 - random errors, 3 terminate")
	numberOfNumbers := flag.Int("numbers", 10*1000*1000, "number of numbers to send")

	flag.Parse()

	tcpAddr, err := net.ResolveTCPAddr("tcp", *interfaceToConnect)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < *numberOfNumbers; i++ {
		number := r1.Intn(100000000)

		var strEcho string

		switch *mode {
		case 1:
			strEcho = fmt.Sprintf("%09d\n", number)
			break
		case 2:
			if number%1613 == 0 {
				strEcho = fmt.Sprintf("O%08d\n", number)
			} else if number%1612 == 0 {
				strEcho = fmt.Sprintf("ololool\n")
			} else {
				strEcho = fmt.Sprintf("%09d\n", number)
			}
			break
		case 3:
			strEcho = fmt.Sprintf("terminate\n")
		}

		writeToServer(conn, strEcho)

		if i%(100*1000) == 0 {
			fmt.Printf("Messages send: %d\n", i)
		}
	}

	conn.Close()
}

func writeToServer(conn *net.TCPConn, message string) {
	_, err := conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Write to server failed:", err.Error())
		os.Exit(1)
	}
}
