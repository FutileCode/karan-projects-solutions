/*
Fibonacci Sequence - Enter a number and
have the program generate the Fibonacci sequence
to that number or to the Nth number.
*/

package main

import (
	"fmt"
	"math/big"
)

func fib(c chan<- *big.Int) {
	x, y := big.NewInt(0), big.NewInt(1)

	for {
		c <- x

		x, y = y, x
		x.Add(x, y)
	}
}

func main() {
	var digits int

	fmt.Println("Calculate Fibonacci sequence up to n. Enter n")
	fmt.Scanln(&digits)

	c := make(chan *big.Int)

	go fib(c)

	for i := 0; i <= digits; i++ {
		fmt.Println(<-c)
	}
}
