package main

import "fmt"

func main() {
	result := calculate(1, 2, sum)
	fmt.Println("sum", result)
	result = calculate(4, 3, sub)
	fmt.Println("sub", result)
	result = calculate(2, 3, multiply)
	fmt.Println("multiply", result)
	result = calculate(4, 2, divide)
	fmt.Println("division", result)
}

func calculate(x, y int, operation func(int, int) int) int {
	result := operation(x, y)

	fmt.Print("x:", x, " y:", y, "\t")

	return result
}

func sum(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
