package main

import "fmt"

func main() {
	name := "Robab"
	age := 56
	weight := 76.5
	height := 165
	woman := true

	fmt.Printf("%s's information\n", name)
	fmt.Println("Woman: ", woman)
	fmt.Println("Age: ", age)
	fmt.Println("Weight: ", weight)
	fmt.Println("height: ", height)

	// for strings
	var strVar string
	strVar = "string values"

	var c rune
	c = '1'

	var number int
	number = 32

	var floatNumber float32
	floatNumber = 3.2

	fmt.Println(strVar, c, number, floatNumber)
}
