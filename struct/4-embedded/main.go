package main

import "fmt"

type Person struct {
	name string
	age  uint
}

type Employee struct {
	Person
	salary uint
}

func main() {
	cj := Person{
		name: "cj",
		age:  35,
	}

	cjEmployee := Employee{
		Person: cj,
		salary: 30_000,
	}

	fmt.Println(cjEmployee.salary)
	fmt.Println(cjEmployee.name)
	fmt.Println(cjEmployee.age)
	fmt.Println(cjEmployee.Person)
}
