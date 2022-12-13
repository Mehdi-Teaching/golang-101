package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	results := make(chan int)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, number := range numbers {
		wg.Add(1)
		go power(number, results)
	}

	go printer(results, &wg)

	wg.Wait()
}

func power(x int, results chan int) {
	<-time.After(1 * time.Second)
	results <- x * x
}

func printer(results chan int, wg *sync.WaitGroup) {
	for result := range results {
		fmt.Println(result)
		wg.Done()
	}
}
