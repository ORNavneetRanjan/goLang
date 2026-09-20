package main

import "fmt"

func calculateArea(width float64, length float64) float64 {
	if width < 0 || length < 0 {
		return -1
	}
	return width * length
}

func fact(n int) int {
	if(n <= 1){
		return 1;
	}
	return n * fact(n-1)
}


func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}


func m1() {
	area := calculateArea(10.45678, 132.04789)
	if area == -1 {
		fmt.Println("Please provide vaild length and widht");
	}else {
		fmt.Printf("Area : %.4f\n", area)
	}

	nums := intSeq()
	fmt.Println(nums())
	fmt.Println(nums())
	fmt.Println(nums())

}
