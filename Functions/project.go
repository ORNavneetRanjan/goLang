package main

import (
	"fmt"
	"strings"
)

type MathError struct {
	Operation string
	InputA    int
	InputB    int
	Message   string
}

const (
	division      = "Division"
	divisionError = "division by zero is not allowed"
)

func (e MathError) Error() string {
	var input []string
	if e.Operation == division {
		input = append(input, fmt.Sprintf("a=%d", e.InputA))
		input = append(input, fmt.Sprintf("b=%d", e.InputB))
	}

	return fmt.Sprintf("Math error is %s (%s): %s",
		e.Operation,
		strings.Join(input, ","),
		e.Message)
}

func Sum(numbers ...int) int{
	defer fmt.Println("Sum finished")
	total := 0

	for _, val := range numbers {
		total += val
	}
	return total
}

func SafeDivision(a, b int) (int, error) {
	if b == 0 {
		return 0, &MathError{
			Operation: division,
			InputA: a,
			InputB: b,
			Message: divisionError,
		}
	}
	return a/b, nil
}


func main() {
	res1, err := SafeDivision(1, 3)
	if err != nil {
			fmt.Println("Error:", err)
			return
	}

	res2, err := SafeDivision(212, 0)
	if err != nil {
			fmt.Println("Error:", err)
			return
	}

	fmt.Println(Sum(res1, res2))
}