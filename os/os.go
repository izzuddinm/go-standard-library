package main

import (
	"fmt"
	"os"
)

func main() {
	// run with go run os.go "production"
	// or
	//  run with go run os.go local development staging production
	args := os.Args

	for index, arg := range args {
		fmt.Println(index, arg)
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error when get hostname os", err.Error())
	} else {
		fmt.Println("Hostname", hostname)
	}
}
