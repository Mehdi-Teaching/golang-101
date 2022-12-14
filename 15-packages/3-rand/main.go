package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	f1 := rand.Float32()
	f2 := rand.Float32()
	fmt.Println(f1, f2)

	key := make([]byte, 10)
	rand.Read(key)

	fmt.Println(key)
}
