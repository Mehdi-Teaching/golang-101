package main

import "fmt"

func main() {
	var weight float32
	var height int

	fmt.Print("enter weight and height in order: ")
	fmt.Scan(&weight, &height)

	heightInMeter := float32(height) / 100

	bmi := weight / (heightInMeter * heightInMeter)

	fmt.Println("Your bmi is ", bmi)

	if bmi < 19 {
		fmt.Println("you are skinny")
	} else if bmi < 25 {
		fmt.Println("you are in perfect condition")
	} else if bmi < 29 {
		fmt.Println("you are overweight")
	} else if bmi < 35 {
		fmt.Println("you are fat")
	} else {
		fmt.Println("you are obese")
	}

}
