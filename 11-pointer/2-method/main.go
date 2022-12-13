package main

import "fmt"

type Holder struct {
	numbers []int
}

func (h Holder) addNumber(x int) {
	h.numbers = append(h.numbers, x)
}

func (h *Holder) addNumberByRef(x int) {
	h.numbers = append(h.numbers, x)
}

func main() {
	h := Holder{
		numbers: make([]int, 0),
	}

	h.addNumber(1)

	fmt.Println(h.numbers)

	h.addNumberByRef(10)

	fmt.Println(h.numbers)
}
