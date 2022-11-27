package main

import (
	"fmt"
)

func main() {
	var age int

	fmt.Print("Your Age: ")
	fmt.Scan(&age)

	fmt.Println("Your age is", age)
}
