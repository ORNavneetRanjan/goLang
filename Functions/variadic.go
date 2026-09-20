package main

import "fmt"

func config(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	total := 0
	for _, val := range arr {
		total += val
	}
	return total
}

func vaild(size int, arr ...int) {
	if size > len(arr) {
		size = len(arr)
	}

	total := 0
	for i := range size {
		total += arr[i]
	}
	fmt.Println(total)
}

func m2() {
	mySlice := make([]int, 10, 15)
	copy(mySlice, []int{12, 12, 13, 42, 452})
	
	fmt.Println(config(mySlice))
	vaild(10)
	vaild(10, mySlice...)

	myArray := [6] int {10, 20, 30, 40, 50, 60}

	vaild(9, myArray[:]...)

}