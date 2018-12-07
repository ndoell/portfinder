package main

import (
	"net"
	"log"
	"fmt"
)


//Get an open port from kernel by using ResolveTCPAddr then test by using ListenTCP.
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
	port, err := GetFreePort()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(port)
}

