package main

import "fmt"

func main() {
	// empty array
	var ids []int

	ids = append(ids, 1)
	ids = append(ids, 2)
	ids = append(ids, 4, 5)

	fmt.Println(ids)
}
