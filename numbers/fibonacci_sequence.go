package main

import (
	"fmt"
)

func main() {
	var digits int
	fmt.Scanln(&digits)

	d1 := 0
	d2 := 1
	next := 0

	for i := 1; i <= digits; i++ {
		if i == 1 {
			fmt.Println(d1)
			continue
		}
		if i == 2 {
			fmt.Println(d2)
			continue
		}

		next = d1 + d2
		d1 = d2
		d2 = next
		fmt.Println(next)
	}
}
