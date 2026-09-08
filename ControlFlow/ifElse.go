package main

import (
	"fmt"
	"math/rand"
)

func m2() {

	tmp := 25

	if tmp > 30 {
		fmt.Println(tmp)
	}else {
		fmt.Printf("Temperateur is high : %d\n", tmp)
	}

	score := rand.Int()
	fmt.Printf("Score : %d\n", score)

	if score >= 90{
		fmt.Println("Grade: A+")
	}else if score >= 80 {
		fmt.Println("Grade: B+")
	}else if score >= 70 {
		fmt.Println("Just passed")
	}else {
		fmt.Println("Bro try harder")
	}

	userAccess := map[string]bool {
		"jane" : true,
		"john" : false,
	}

	if hasAccess, ok :=  userAccess["john"] ; ok && hasAccess {
		fmt.Println("Jane can access the system")
	}else {
		fmt.Println("Jane dont have access to this")
	}

}