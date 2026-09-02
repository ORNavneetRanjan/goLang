package main

import (
	"fmt"
	"strings"
)

func m1() {

	const HOST string = "localHost"

	

	var getting string 
	getting = "this is it"
	fmt.Println(getting)

	var isRunning bool
	isRunning = true
	fmt.Println(isRunning)

	var firstName , lastName string
	firstName = "Navneet"
	lastName = "Ranjan"

	fmt.Println(firstName + " " + lastName)

	email := "kumar@gmail.com"
	var address []string = strings.Split(email, ".")
	fmt.Println(address, email)

	var n = 10
	fmt.Println(n * 10)





	// fmt.Println("hello")
	// fmt.Printf("This is it")
	// fmt.Println(1 + 2)
	// fmt.Println(true)

	// fmt.Printf("%+v\n", []int{1, 2, 3})
	// var t any
	// fmt.Printf("%+v\n", t)

}
