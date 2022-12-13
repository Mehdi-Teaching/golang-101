package main

import (
	"fmt"
	"math/rand"
)

func main() {

	if ok := randBool(); ok {
		fmt.Println("it is ok")
	}

	// ok is not accessible here!
	// fmt.Println(ok)
}

func randBool() bool {
	return rand.Float32() > 0.5
}
