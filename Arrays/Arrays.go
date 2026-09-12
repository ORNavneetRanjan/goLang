package main

import "fmt"

func m1() {

	var number [2]int

	fmt.Println(number)

	var matrix [2][3] int;
	matrix[0][0] = 1;
	matrix[1][2] = 3;
	
	for _, val := range matrix {
		for _, val1 := range val {
			fmt.Println(val1)
		}
	}

	fmt.Printf("%+v\n", matrix)

	arr := [] string {"GO", "Python", "Rust"}
	fmt.Println(arr)

}