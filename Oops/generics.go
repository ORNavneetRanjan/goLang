package main

import "fmt"

type Numbers interface {
	int | int8 | int16 | float32 | float64
}

func Sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func SumAll[T Numbers](numbers ...T) T {
	var total T
	total = 0
	for _, number := range numbers {
		total += number
	}
	return total 
} 

func main4(){
	grades := []int{90, 85}
	people := []string{"Jane", "Jhon", "Mark"}

	fmt.Println(len(grades), len(people), Sum(grades...))
	fmt.Printf("%T\n", SumAll(1, 2, 24, 5))
}
