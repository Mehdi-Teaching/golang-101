package main

import "fmt"

func main() {
	names := []string{"kennedy", "johnson", "nixon", "ford"}

	for i := 0; i < len(names); i++ {
		name := names[i]
		fmt.Println(i, name)
	}

	fmt.Println("===========================")

	j := 0
	for j < len(names) {
		name := names[j]
		fmt.Println(j, name)
		j++
	}

}
