package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Muhammad Ayom Izzuddin", "Izzuddin"))
	fmt.Println(strings.Split("Muhammad Ayom Izzuddin", " "))
	fmt.Println(strings.ToLower("Muhammad Ayom Izzuddin"))
	fmt.Println(strings.ToUpper("Muhammad Ayom Izzuddin"))
	fmt.Println(strings.Trim("     Muhammad Ayom Izzuddin     ", " "))
	fmt.Println(strings.ReplaceAll("Muhammad John Ayom Izzuddin John", "John", ""))
}
