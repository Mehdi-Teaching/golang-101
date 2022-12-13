package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {

	x := 2
	fmt.Println("x is", x)

	x, err := power(x)
	if err != nil {
		panic(err)
	}

	fmt.Println("result is", x)

}

func power(x int) (int, error) {
	if rand.Float32() > 0.5 {
		return 0, errors.New("failed to calculate")
	} else {
		return x * x, nil
	}
}
