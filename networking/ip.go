package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-address\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]

	address := net.ParseIP(name)

	if address == nil {
		fmt.Println("Invalid Address")
	} else {
		fmt.Println("The address is ", address)
	}

	os.Exit(0)
}
