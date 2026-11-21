package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
	JoindAt  	time.Time
}

func NewEmployee(id int,
	firstName, lastName, position string, isActive bool) Employee {

	return Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Position:  position,
		IsActive:  isActive,
		JoindAt:   time.Now(),
	}

}

func (e *Employee) FullName(f Employee) string {
	return f.FirstName + " " + f.LastName
}

func (e Employee) Deactivate() {
	e.IsActive = false
	fmt.Println(e)
}

func m1() {
	
	navneet := Employee{
		ID: 1,
		FirstName: "Navneet",
		LastName: "Ranjan",
		Position: "Night",
		Salary: 	1000,
		IsActive: true,
		JoindAt: time.Now(),	
	}

	fmt.Printf("%+v \n", navneet)

	ranjan := NewEmployee(2, "Ranjan", "Kumar", "Day", false)

	fmt.Printf("%+v \n",ranjan)

	ranjan = navneet
	fmt.Println(ranjan)
	ranjan.Salary = 1000000
	fmt.Println(navneet)

	ptr_ranjan := &navneet
	ptr_ranjan.Salary = 20000
	fmt.Println(navneet)

}