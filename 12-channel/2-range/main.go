package main

import (
	"fmt"
	"time"
)

func main() {
	results := make(chan int)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, number := range numbers {
		go power(number, results)
	}

	go printer(results)

	<-time.After(3 * time.Second)

	fmt.Println("finished :)")
}

func power(x int, results chan int) {
	<-time.After(1 * time.Second)
	results <- x * x
}

func printer(results chan int) {
	for result := range results {
		fmt.Println(result)
	}
}
