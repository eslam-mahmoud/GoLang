//https://tour.golang.org/moretypes/26
//Exercise: Fibonacci closure

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var step int
	var number int
	var previousNumber int
	return func() int {
		if step == 0 {
			step++
		} else if step == 1 {
			step++
			number += 1
		} else {
			previousNumber, number = number, number + previousNumber
		}
		return number
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
