package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func main() {
	target := "scanme.nmap.org"
	// port := 80
	// portStr := strconv.Itoa(port)
	// address := net.JoinHostPort(target, portStr)

	// set a timeout
	dialer := net.Dialer{
		Timeout: 5 * time.Second,
	}

	maxRetries := 3

	// ONE CONNECTION
	// conn, err := dialer.Dial("tcp", address)
	// if err != nil {
	// 	log.Fatalf("unable to connect to %s: %v", address, err)
	// }
	// defer conn.Close()

	// fmt.Printf("connection to %s was successful\n", address)

	var wg sync.WaitGroup

	for port := 80; port <= 85; port++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			address := net.JoinHostPort(target, strconv.Itoa(port))

			fmt.Printf("[PORT %d - %s]\n", port, address)

			for i := 0; i < maxRetries; i++ {
				conn, err := dialer.Dial("tcp", address)
				if err == nil {
					defer conn.Close()
					fmt.Printf("[SUCCESS] %s\n", address)
					break
				}

				backoff := time.Duration(1<<i) * time.Second
				fmt.Printf("[%d] Attempt %d failed. Waiting %v...\n", port, i+1, backoff)
				time.Sleep(backoff)
			}
		}()

	}

	wg.Wait()
}
