package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]

	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The Ip address is : ", addr.String())
	}
	os.Exit(0)
}

// Ipv4
// go run Ipcheck.go 127.0.0.1
// The Ip address is : 127.0.0.1

// Ipv6
// go run Ipcheck.go 0.0.0.0.0.0.0.1
// The Ip address is : ::1
