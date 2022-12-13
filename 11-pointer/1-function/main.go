package main

import "fmt"

type Person struct {
	Name string
}

func main() {

	fmt.Println("============ set variable ============")

	t := 0
	assignNumber(t)
	fmt.Println("passed by value:", t)

	assignNumberByRef(&t)
	fmt.Println("passed by reference:", t)

	fmt.Println("============ update struct ============")

	person := Person{
		Name: "name",
	}

	updateStruct(person)
	fmt.Println("passed by value:", person)

	updateStructByRef(&person)
	fmt.Println("passed by reference:", person)

	fmt.Println("============ append array ============")

	var d []int

	d = append(d, 0)

	assignArray(d)
	fmt.Println("passed by value:", d)

	assignArrayByRef(&d)
	fmt.Println("passed by reference:", d)

	fmt.Println("============ update variable ============")

	p := []Person{
		{Name: "name"},
	}

	updateArrayElement(p)
	fmt.Println("passed by value:", p)

	s := []Person{
		{Name: "name"},
	}

	updateArrayElementByReference(&s)
	fmt.Println("passed by reference:", s)
}

func assignNumber(x int) {
	x = 2
}

func assignNumberByRef(x *int) {
	*x = 2
}

func updateStruct(p Person) {
	p.Name = "update_struct"
}

func updateStructByRef(p *Person) {
	(*p).Name = "update_field"
}

func assignArray(x []int) {
	x[0] = 1000
	x = append(x, 2)
}

func assignArrayByRef(x *[]int) {
	*x = append(*x, 2)
}

func updateArrayElement(x []Person) {
	x[0].Name = "updated"
}

func updateArrayElementByReference(x *[]Person) {
	(*x)[0].Name = "updated"
}
