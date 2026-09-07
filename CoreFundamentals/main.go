package main

import "fmt"

type LogLevel int

const (
	LevelTrace LogLevel = iota + 2
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
)

var levelName = []string{"Trace", "Debug", "Info", "Warning", "Error"}

func (l LogLevel) func1() string {
	if(l < 0 || l >= LogLevel((len(levelName)))) {return "Unknow" }
	return levelName[l];
}

func printLogLevel(level LogLevel) {
	fmt.Printf("Log Level: %d %s\n", level, level.func1())
}

func main() {
	printLogLevel(LevelTrace)
	printLogLevel(LevelDebug)
	printLogLevel(LevelInfo)
	printLogLevel(LevelWarning)
	printLogLevel(LevelError)

	printLogLevel(10)
	printLogLevel(1067890789098767)
}