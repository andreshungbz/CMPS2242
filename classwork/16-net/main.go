package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func main() {
	dialer := net.Dialer{Timeout: 5 * time.Second}
	target := "scanme.nmap.org"
	maxRetries := 3
	ports := 512
	var wg sync.WaitGroup

	for p := 1; p <= ports; p++ {
		port := strconv.Itoa(p)
		address := net.JoinHostPort(target, strconv.Itoa(p))
		wg.Add(1)

		go func() {
			defer wg.Done()
			fmt.Printf("[SCAN %s PORT %s]\n", address, port)

			for i := range maxRetries {
				conn, err := dialer.Dial("tcp", address)
				if err == nil {
					defer conn.Close()
					fmt.Printf("[SUCCESS] %s\n", address)
					break
				}

				backoff := time.Duration(1<<i) * time.Second
				fmt.Printf("[PORT %s ATTEMPT %d FAIL] Waiting %v...\n", port, i+1, backoff)
				time.Sleep(backoff)
			}
		}()
	}

	wg.Wait()
}
