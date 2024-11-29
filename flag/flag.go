package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "root", "database password")
	var host *string = flag.String("host", "localhost", "database host")
	var port *int = flag.Int("port", 8080, "database port")

	flag.Parse()

	fmt.Println("Username", *username)
	fmt.Println("Password", *password)
	fmt.Println("Host", *host)
	fmt.Println("Port", *port)

	// How to run

	// 1. PS D:\Ayom\Exercise\go-standard-library\flag> go run flag.go
	// Username root
	// Password root
	// Host localhost
	// Port 8080

	// 2. PS D:\Ayom\Exercise\go-standard-library\flag> go run flag.go -username="nexus_app" -password="nexus_secret" -host="192.168.10.42" -port=5027
	// Username nexus_app
	// Password nexus_secret
	// Host 192.168.10.42
	// Port 5027
}
