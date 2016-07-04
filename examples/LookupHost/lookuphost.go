package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage : go run lookuphost.go ")
		os.Exit(1)
	}

	name := os.Args[1]

	addrs, err := net.LookupHost(name)

	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}

	for _, s := range addrs {
		fmt.Println(s)
	}

	os.Exit(0)
}

// returns strings  not ip address values

// go run lookuphost.go google.com
// 216.58.197.78
// 2404:6800:4007:800::200e

// go run lookuphost.go google.co.in
// 216.58.197.67
// 2404:6800:4007:800::2003
