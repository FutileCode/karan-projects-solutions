/*
Develop a converter to convert a decimal
number to binary or a binary number to its
decimal equivalent.
*/

package main

import (
	"flag"
	"fmt"
)

func main() {
	conversion := flag.Int("conversion", 1, "1 = dec -> bin | 2 = bin -> dec")
	number := flag.Int("number", 0, "bin or dec number")
	flag.Parse()

	switch *conversion {
	case 1:
		fmt.Printf("%d as bin -> %b \n", *number, *number)
	case 2:
		fmt.Printf("%d as dec -> %b \n", *number, *number)
	default:
		fmt.Println("conversion flag must be 1 or 2")
	}
}
