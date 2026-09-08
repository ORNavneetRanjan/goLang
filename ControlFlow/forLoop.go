package main

import "fmt"

func m1() {
	// only for loop is avilable in go

	for i := 1; i<=10;i++ {
		fmt.Println(i)
	}

	//while style for loop
	k := 10
	for k > 0 {
		k--
	}
	fmt.Println(k)

	//counter 
	counter := 0
	for {
		fmt.Println(counter)
		counter++
		if(counter >= 5){
			break
		}
	}

	//skiping 

	for i := 1; i<=10;i++ {
		if(i % 2 == 0){
			continue
		}
		fmt.Println(i)
	}

	var languages = []string  {"Go", "Python", "Java", "C++"}

	for index,val := range languages {
		fmt.Printf("index: %d => value : %s\n", index, val)
	}

	for _, val := range languages {
		fmt.Println(val)
	}

	for index := range languages {
		fmt.Println(index)
	}
}