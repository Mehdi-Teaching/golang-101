package main

import (
	"fmt"
	"time"
)

func main() {

	results := make(chan int)

	fmt.Println("before goroutines")

	go sumArray([]int{1, 2, 3, 4, 5}, results)
	go sumArray([]int{6, 7, 8, 9, 10}, results)

	fmt.Println("after goroutines in main")

	x := <-results
	fmt.Println("sum:", x)

	x = <-results

	fmt.Println("sum:", x)

}

func sumArray(numbers []int, result chan int) {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	// delay
	<-time.After(1 * time.Second)

	result <- sum
}
