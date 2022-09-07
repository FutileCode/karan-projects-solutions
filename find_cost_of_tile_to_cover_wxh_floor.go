/*
Find Cost of Tile to Cover W x H
Floor - Calculate the total cost of tile
it would take to cover a floor plan
of width and height, using a cost entered by the user.
*/

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
