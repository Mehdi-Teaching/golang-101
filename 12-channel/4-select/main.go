package main

import (
	"fmt"
	"time"
)

func main() {

	result := make(chan int)

	go complexCalculation(result)

	select {
	case x := <-result:
		fmt.Println("result is", x)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
	}

}

func complexCalculation(result chan int) {
	<-time.After(10 * time.Second)
	result <- 1000
}
