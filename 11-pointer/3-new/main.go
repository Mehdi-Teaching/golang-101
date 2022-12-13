package main

import (
	"fmt"
	"strings"
)

func main() {
	x := new(int)

	*x = 2

	fmt.Println(*x)

	appender := new(strings.Builder)
	appender.WriteString("hello ")
	appender.WriteString("world")
	fmt.Println(appender.String())
}
