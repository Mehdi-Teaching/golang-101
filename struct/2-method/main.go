package main

import "fmt"

type Person struct {
	name string
	age  uint
}

func (p Person) sayMyName() {
	fmt.Println("- you are", p.name)
	fmt.Println("+ You're god damn right")
}

func main() {
	heisenberg := Person{
		name: "heisenberg",
		age:  50,
	}

	heisenberg.sayMyName()
}
