package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error.")
	NotFoundError   = errors.New("not found error.")
)

func GetById(name string) (string, error) {
	if name == "" {
		return "", ValidationError
	} else if name == "john" {
		return "", NotFoundError
	} else {
		return name, nil
	}
}

func main() {
	namesToTest := []string{"john", "", "doe"}

	for _, name := range namesToTest {
		result, err := GetById(name)
		if err != nil {
			switch {
			case errors.Is(err, ValidationError):
				fmt.Println("Validation Error:", err.Error())
			case errors.Is(err, NotFoundError):
				fmt.Println("Not Found Error:", err.Error())
			default:
				fmt.Println("Unknown Error:", err.Error())
			}
		} else {
			fmt.Println("Retrieved:", result)
		}
	}
}
