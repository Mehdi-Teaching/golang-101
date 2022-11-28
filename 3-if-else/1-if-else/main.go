package main

import "fmt"

func main() {
	var age int

	fmt.Print("What is your Age? ")
	fmt.Scan(&age)

	if age < 18 {
		fmt.Println("You are not eligible for driving")
	} else {
		fmt.Println("You can drive ;)")
	}
}
