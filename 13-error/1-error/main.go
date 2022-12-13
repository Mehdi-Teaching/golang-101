package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {

	x, err := power(10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("the result is", x)
	}

}

func power(x int) (int, error) {
	if rand.Float32() > 0.5 {
		return 0, errors.New("failed to calculate")
	} else {
		return x * x, nil
	}
}
