package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// array[inclusive-index:exclusive-index]
	fmt.Println("numbers ", numbers)
	fmt.Println("numbers[:5] ", numbers[:5])
	fmt.Println("numbers[2:] ", numbers[2:])
	fmt.Println("numbers[2:5] ", numbers[2:5])

	// reference to numbers
	slice := numbers[:3]

	slice[1] = 999

	fmt.Println("slice: ", slice)
	fmt.Println("numbers: ", numbers)

}
