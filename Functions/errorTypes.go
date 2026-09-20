package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

var DivideByZero error = errors.New("Cannot divide by Zero")
var NumberTooBig error = errors.New("Number too big")

type OpError struct {
	Op   		string
	code 		int
	message string
	time 		time.Time
}

func(op OpError) Error() string {
	return op.message
}

func NewOpError(Op string, code int, message string, time time.Time) *OpError{
	return &OpError{
		Op : Op,
		code: code,
		message: message,
		time: time,
	}
}

func DoSomething() error {
	return NewOpError("doSomething", 100, "do Something failed", time.Now())
}


func checkDivisiblity(a, b int) (int, error) {
	if(b == 0) {return -1, DivideByZero}
	if (a > math.MaxInt / 2 || b > math.MaxInt / 2) {
		return math.MaxInt, NumberTooBig
	}
	return a/b, nil
}

func display(a int, e error){
	if(e != nil){
		fmt.Println(e);
		return 
	}
	fmt.Println(a)
}

func m4() {
	display(checkDivisiblity(10, 3))
	display(checkDivisiblity(10, 0))
	display(checkDivisiblity(math.MaxInt, 10))

}