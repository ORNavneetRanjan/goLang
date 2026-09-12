package main

import "fmt"


func modifyValue(val int){
	val *= 10
	fmt.Printf("Modified val: %d\n",val)
}

func byAdd(val *int) {
	if val == nil {
		fmt.Println(`val is nil`)
		return 
	}
	*val *= 100
}

func main() {

	age := 10
	fmt.Printf("age: %d\n", &age)
	add := &age
	modifyValue(age)
	byAdd(&age)
	fmt.Printf("age: %d\n", *add)

}