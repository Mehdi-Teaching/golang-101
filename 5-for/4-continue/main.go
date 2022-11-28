package main

import "fmt"

func main() {
	length := 5
	names := []string{"kennedy", "johnson", "nixon", "ford"}

	fmt.Println("names with length smaller equal than", length)

	for i, name := range names {
		if len(name) <= length {
			fmt.Println("index", i, "name", name)
		}
	}

	fmt.Println("===============================")

	for i, name := range names {
		if len(name) > length {
			continue
		}

		fmt.Println("index", i, "name", name)
	}
}
