// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage:
// Navigate to bhg-scanner/main and run `go build`, then execute ./main

package scanner

import (
	"fmt"
	"net"
	"sort"
	"time"
)

type port struct {
	portNumber int
	status PortStatus
}

func worker(ports chan int, results chan port, address string) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", address, p)    
		conn, err := net.DialTimeout("tcp", address, 1 * time.Second )
		if err != nil { 
			results <- port{p,1}
			continue
		}
		conn.Close()
		results <- port{p,2}
	}
}

func PortScanner(address string, portsToScan []int) (int, int) {  
	openPorts := make([]int, 0)
	closedPorts := make([]int, 0)

	ports := make(chan int, len(portsToScan))
	results := make(chan port)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, address)
	}

	go func() {
		for i := 1; i <= len(portsToScan); i++ {
			ports <- portsToScan[i - 1]
		}
	}()

	for i := 0; i < len(portsToScan); i++ {
		p := <- results
		switch p.status {
		case Closed:
			closedPorts = append(closedPorts, p.portNumber)
		case Open:
			openPorts = append(openPorts, p.portNumber)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openPorts)
	sort.Ints(closedPorts)

	fmt.Printf("\nOpen Ports\n")
	for _, p := range openPorts {
		fmt.Printf(", %d", p)
	}
	
	fmt.Printf("\nClosed Ports\n")
	for _, p := range closedPorts {
		fmt.Printf(", %d", p)
	}

	return len(openPorts), len(closedPorts)
}
