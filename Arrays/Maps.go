package main

import (
	"fmt"
)

func m3() {

	studentGrades := map[string]int{
		"Navneet": 20,
		"Ranjan":  17,
		"Kumar":   13,
	}

	fmt.Println(studentGrades)
	studentGrades["Hello"] = 90
	fmt.Println(studentGrades)

	navneet, ok := studentGrades["Navneet"]
	delete(studentGrades, "Navneet")

	fmt.Println(navneet, ok)
	fmt.Printf("Score: %d\n", studentGrades["random"])

	configs := map[string]int{};
	fmt.Println(configs)

	var maps map[string]int;
	fmt.Println(maps)


}