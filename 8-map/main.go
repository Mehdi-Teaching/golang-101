package main

import "fmt"

func main() {
	var people map[uint]string

	people = make(map[uint]string)

	people[0] = "Mehdi"
	people[1] = "Ali"
	people[2] = "Hosna"
	people[3] = "Qazal"

	printPeople(people)

	delete(people, 2)

	printPeople(people)

	fmt.Println("len:", len(people))

	person := people[0]
	printPerson(0, person)

	person, ok := people[2]
	if ok {
		printPerson(2, person)
	} else {
		fmt.Printf("person with %d id is not avilable\n", 2)
	}

}

func printPeople(people map[uint]string) {
	for key, value := range people {
		fmt.Printf("id: %d , name: %s\n", key, value)
	}
	fmt.Println("===================")
}

func printPerson(id uint, name string) {
	fmt.Println("id:", id, " name:", name)
}
