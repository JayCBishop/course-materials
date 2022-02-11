package main

import "bhg-scanner/scanner"

func main(){
	var portsToScan []int
	for i := 1; i <= 1024; i++ {
		portsToScan = append(portsToScan, i)
	}
	scanner.PortScanner("scanme.nmap.org", portsToScan)
}