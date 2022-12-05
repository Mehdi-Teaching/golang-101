package main

import "fmt"

func main() {
	// searching for a name
	wanted := "johnson"
	names := []string{"kennedy", "johnson", "nixon", "ford"}

	index := -1
	for i, name := range names {
		if name == wanted {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("didn't found him :(")
	} else {
		fmt.Println("found him! :)")
		fmt.Printf("%s is at room %d\n", wanted, index)
	}
}
