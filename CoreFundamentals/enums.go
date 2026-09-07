package main

import "fmt"


const (
	Sunday = iota + 2
	Monday
	Tuesday
	Wednesday
)

type level int

const (
	LogError level = iota
	LogWarn
	LogInfo
	LogDebug
)


func m2() {
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)

	fmt.Println(LogError)
}