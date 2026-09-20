package main

import (
	"errors"
	"fmt"
	"strings"
)

func divide(a, b int) (int, error, int) {
	if b == 0 {
		return 10, errors.New("Divide by zero"), 0
	}
	return a/b, nil, a%b
}

func splitNames(fullName string) (firstName, lastName string) {
	parts := strings.Split(fullName, " ")
	firstName = parts[0]
	lastName = parts[1]
	return 
}

func m3() {
	fmt.Println(divide(10, 2))
	fmt.Println(divide(10, 3))
	fmt.Println(divide(10, 0))

	fmt.Println(splitNames("Navneet Ranjan"))
	fmt.Println(splitNames("Navneet Ranjan Kumar"))
}