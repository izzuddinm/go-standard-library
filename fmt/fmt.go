package main

import "fmt"

func main() {
	firstName := "Muhammad"
	middleName := "Ayom"
	lastName := "Izzuddin"

	fmt.Println("Hello '", firstName, middleName, lastName, "'")

	fmt.Printf("Hello '%s' %s %s\n", firstName, middleName, lastName)
}
