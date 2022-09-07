/*
Find PI to the Nth Digit - Enter a number and have the
program generate PI up to that many decimal places.
Keep a limit to how far the program will go.
*/

package main

import (
	"fmt"
	"math"
)

func main() {

	var digits int

	for {
		fmt.Println("Enter number of digits to print pi to (max 48)")
		fmt.Scanln(&digits)

		if digits > 48 {
			fmt.Println("!!! Max digits of 48 !!!")
		} else {
			break
		}
	}

	pi := fmt.Sprintf("%.[2]*[1]f", math.Pi, digits)
	fmt.Printf("Pi to %v digits is:\n%v", digits, pi)

}
