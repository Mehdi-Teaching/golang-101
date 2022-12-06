package main

import "fmt"

func main() {

	fmt.Println(sum(1, 2, 3, 4, 5))

	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(sum(numbers...))
}

func sum(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}

	return result
}
