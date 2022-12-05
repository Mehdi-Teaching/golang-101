package main

import "fmt"

func main() {
	x := 10
	y := 5

	xySum := sum(x, y)
	xySub := sub(x, y)

	fmt.Println("sum:", xySum)
	fmt.Println("sub:", xySub)
}

func sum(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}
