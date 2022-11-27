package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("start typing. I will echo everything you write back to you :)")

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("you typed: ", text)
	}
}
