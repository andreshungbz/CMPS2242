package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	target := "scanme.nmap.org"
	port := 80
	address := net.JoinHostPort(target, fmt.Sprintf("%d", port))

	// set a timeout
	dialer := net.Dialer{
		Timeout: 5 * time.Second,
	}

	conn, err := dialer.Dial("tcp", address)
	if err != nil {
		log.Fatalf("unable to connect to %s: %v", address, err)
	}
	defer conn.Close()

	fmt.Printf("connection to %s was successful\n", address)
}
