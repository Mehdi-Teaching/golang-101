package main

import "fmt"

type Person struct {
	name string
	age  uint
}

func main() {
	mehdi := Person{
		name: "Mehdi",
		age:  30,
	}

	fmt.Println("name:", mehdi.name)
	fmt.Println("age:", mehdi.age)
	fmt.Println(mehdi)
	fmt.Printf("full info: %+v", mehdi)
}
