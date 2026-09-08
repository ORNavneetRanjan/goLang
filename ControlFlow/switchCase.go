package main

import (
	"fmt"
	"math/rand"
	"time"
)

func m3() {

	day := "Sunday"
	fmt.Println("Today is ", day)

	switch day {
		case "Sunday", "Saturday":
			fmt.Println("Holiday")
		case "Monday", "Tuesday" :
			fmt.Println("Too many meetings")
		default:
			fmt.Println("Work and go home early")

	}

	switch hour := time.Now().Hour();  {
		case hour <= 12:
			fmt.Print("Good Morning")
		case hour < 17:
			fmt.Println("Good Afternoon")
		default:
			fmt.Println("Just sleep bro: ", hour)
	}

	checkType := func(i any){
		switch v := i.(type) {
		case int:
			fmt.Printf("Integer : %d\n", v)
		case string:
			fmt.Printf("String: %s\n", v)
		case bool:
			fmt.Printf("Boolean: %t\n",v)
		default:
			fmt.Printf("Unknown type: %T\n", v)
		
		}
	}

	checkType(10)
	checkType("10")
	checkType(10 == rand.Int())
	checkType(10.2340)
	checkType(func ()  {})
	checkType(checkType)

}