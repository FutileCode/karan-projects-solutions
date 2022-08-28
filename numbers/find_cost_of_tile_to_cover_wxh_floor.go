package main

import (
	"fmt"
)

func main() {
	var w, h, p float32

	fmt.Println("Enter width (m)")
	fmt.Scanln(&w)
	fmt.Println("Enter height (m)")
	fmt.Scanln(&h)
	fmt.Println("Enter price per metre squared")
	fmt.Scanln(&p)

	fmt.Printf("Price for floor is %.2f", w*h*p)
}
