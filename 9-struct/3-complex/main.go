package main

import "fmt"

type Person struct {
	name string
	age  uint
}

type Employee struct {
	person Person
	salary uint
}

func main() {
	cj := Person{
		name: "cj",
		age:  35,
	}
	macDonaldEmployee := Employee{
		person: cj,
		salary: 30_000,
	}

	fmt.Println(macDonaldEmployee.salary)
	fmt.Println(macDonaldEmployee.person)
}
