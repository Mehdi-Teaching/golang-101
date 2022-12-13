package main

import (
	"fmt"
	"math/rand"
)

func main() {

	value := randomValue()
	switch value.(type) {
	case bool:
		fmt.Println("bool")
	case int:
		fmt.Println("int")
	case float32:
		fmt.Println("float32")
	case string:
		fmt.Println("string")
	}
}

func randomValue() any {
	r := rand.Float32()

	switch {
	case r < 0.3:
		return 10
	case r < 0.6:
		return "hello"
	case r < 0.9:
		return true
	default:
		return 2.2
	}
}
