package main

import (
	"fmt"
	"slices"
)

func m2() {
	names := []string{"Navneet", "Ranjan", "Kumar"}

	fmt.Println(names)

	items := make([]int, 3, 10)
	fmt.Println(len(items), cap(items))

	items = items[:5]
	fmt.Println(len(items), cap(items))

	items = items[5:]
	fmt.Println(len(items), cap(items))
	
	for i := 10;i<20;i++ {
		items = append(items, i)
	}
	fmt.Println(len(items), cap(items))
	fmt.Println(items)

	if slices.Contains(items, 12) {
		fmt.Println(`12 is their `)
	}
}