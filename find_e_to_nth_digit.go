/*
Find e to the Nth Digit - Just like the previous
problem, but with e instead of PI. Enter a number and
have the program generate e up to that many decimal
places. Keep a limit to how far the program will go.
*/

package main

import (
	"fmt"
	"math"
)

func main() {

	var digits int

	for {
		fmt.Println("Enter number of digits to print E to (max 48)")
		fmt.Scanln(&digits)

		if digits > 48 {
			fmt.Println("!!! Max digits of 48 !!!")
		} else {
			break
		}
	}

	e := fmt.Sprintf("%.[2]*[1]f", math.E, digits)
	fmt.Printf("E to %v digits is:\n%v", digits, e)

}
