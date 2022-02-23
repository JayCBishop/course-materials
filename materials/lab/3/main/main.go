// Build and Use this File to interact with the shodan package
// In this directory lab/3/shodan/main:
// go build main.go
// SHODAN_API_KEY=YOURAPIKEYHERE ./main <search term>

package main

import (
	"fmt"
	"log"
	"os"
	"encoding/json"
	"shodan/shodan"
)

func main() {
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)

	if len(os.Args) < 2 {
		myip, err := s.Utility()
		if err != nil {
			log.Panicln(err)
		}

		fmt.Printf("Your IP Address is: %s\n", *myip)

		host, err := s.SearchByIp(*myip)
		if err != nil {
			log.Panicln(err)
		}

		h, _ := json.Marshal(host)
		fmt.Println(string(h))

		return
	}

	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf(
		"Query Credits: %d\nScan Credits:  %d\n\n",
		info.QueryCredits,
		info.ScanCredits)

	hostSearch, err := s.HostSearch(os.Args[1])

	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("Host Data Dump\n")
	for _, host := range hostSearch.Matches {
		fmt.Println("==== start ",host.IPString,"====")
		h, _ := json.Marshal(host)
		fmt.Println(string(h))
		fmt.Println("==== end ",host.IPString,"====")
	}

	fmt.Printf("IP, Port\n")
	for _, host := range hostSearch.Matches {
		fmt.Printf("%s, %d\n", host.IPString, host.Port)
	}
}
