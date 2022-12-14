package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	formatted := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(formatted)

	// time.Kitchen
	// time.ANSIC
	formatted = time.Now().Format(time.ANSIC)
	fmt.Println(formatted)

	add := time.Now().Add(24 * time.Hour)
	fmt.Println(add.Format(time.RubyDate))

	fmt.Println("Sleeping...")
	time.Sleep(1 * time.Second)
	fmt.Println("I am awake :)")

}
