package main

import "fmt"

func main() {
	fmt.Println(factor(3))
	fmt.Println(factor(5))
}

func factor(n uint) uint {
	if n <= 1 {
		return 1
	}

	return n * factor(n-1)
}
