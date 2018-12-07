package main

import (
//	"math/rand"
//	"time"
	"net"
	"log"
	"fmt"
)

//func randport(min, max int) int {
//	return rand.Intn(max - min) + min
//}

//Check if a port is open by connecting.
func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func main() {
//	rand.Seed(time.Now().UTC().UnixNano())
//	GetFreePort(randport(5000, 9000))

	port, err := GetFreePort()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(port)
}

