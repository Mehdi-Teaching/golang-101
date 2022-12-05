package main

import "fmt"

func main() {
	var bmi uint

	fmt.Print("your bmi: ")
	fmt.Scan(&bmi)

	switch {
	case bmi < 18:
		fmt.Println("you are skinny")
	case bmi < 25:
		fmt.Println("you have a perfect condition")
	case bmi < 30:
		fmt.Println("you are overweight")
	default:
		fmt.Println("you are fat")
	}
}
