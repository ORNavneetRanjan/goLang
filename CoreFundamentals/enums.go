package main

import "fmt"


const (
	Sunday = iota + 2
	Monday
	Tuesday
	Wednesday
)

type LogLevel int

const (
	LogError LogLevel = iota
	LogWarn
	LogInfo
	LogDebug
)


func main() {
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)

	fmt.Println(LogError)
}