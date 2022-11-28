package main

import "fmt"

func main() {
	var size int

	fmt.Print("How many numbers: ")
	fmt.Scan(&size)

	if size < 0 {
		fmt.Println("size must be positive")
	}

	numbers := make([]int, size)

	// scan numbers in a loop
	// in the following lessons

	fmt.Println(numbers)
}
