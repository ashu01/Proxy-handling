package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		// fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		fmt.Println("Usage : go run resolveip.go <hostname> ")
		os.Exit(1)
	}

	name := os.Args[1]

	addr, err := net.ResolveIPAddr("ip", name)

	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}

	fmt.Println("Resolved address is ", addr.String())
	os.Exit(0)
}

// go run resolveip.go ashutoshkg.mutohq.com
// Resolved address is  188.166.255.177
