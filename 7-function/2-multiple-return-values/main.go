package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	count, sum := countSum(numbers)
	fmt.Println(numbers)
	fmt.Println("count", count)
	fmt.Println("sum", sum)

}

func countSum(numbers []int) (int, int) {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	length := len(numbers)

	return length, sum
}
