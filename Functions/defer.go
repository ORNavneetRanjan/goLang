package main

import "fmt"

func simpleDefer() {
	fmt.Println("Function simpleDefer: Start")
		defer fmt.Println("Function simpleDefer: End")
	fmt.Println("Function simpleDefer: Middel")
}

func displayDefer() {
	fmt.Println("Inside the display function")
	simpleDefer()
	defer fmt.Println("Ones this is over call me")
	defer fmt.Println("Close the files of display function")
	simpleDefer()
	fmt.Println("Display function Ends Here")
}
