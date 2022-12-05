package main

import "fmt"

func main() {
	var age uint

	fmt.Print("Your Age: ")
	fmt.Scan(&age)

	legal := age >= 18

	switch legal {
	case true:
		fmt.Println("you are at legal age")
	case false:
		fmt.Println("you are not at legal age")
	}
}
