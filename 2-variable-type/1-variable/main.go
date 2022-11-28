package main

import "fmt"

func main() {
	// you can define variable this way
	// named typed
	var x int
	x = 2

	// or this way.
	// unnamed type
	// short variable declaration
	y := 3

	z := x + y

	fmt.Println("x is ", x)
	fmt.Println("y is ", y)
	fmt.Println("x + y = ", z)
}
