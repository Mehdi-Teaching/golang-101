package main

import "fmt"

func main() {
	names := []string{"kennedy", "johnson", "nixon", "ford"}

	for i := range names {
		name := names[i]
		fmt.Println(i, name)
	}

	fmt.Println("==================================")

	for i, name := range names {
		fmt.Println(i, name)
	}
}
