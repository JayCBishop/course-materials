# README

## Building

In the `lab/3/shodan/main` directory, run `go build main.go`.

## Running

This program currently can be ran in two ways.

### Searching Shodan

This program still functions as it did prior to my changes when an argument is passed to it. To search shodan using the `shodan/host/search` method, the program can be ran as follows:

`SHODAN_API_KEY=YOURAPIKEYHERE ./main <search term>`

### Retrieving Services For Current IP 

When running this program without the search term argument, instead of panicing, the program will now retrieve the client's IP through Shodan's `/tools/myip` method. The program will then query `shodan/host/{ip}` with the client's ip, outputting all services on the ip. Here is an example command-line usage:

`SHODAN_API_KEY=YOURAPIKEYHERE ./main`
