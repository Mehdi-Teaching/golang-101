package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "mehdi"
	result := strings.Repeat(name, 5)
	output("repeat", name, result)

	output("length", name, len(name))

	name = "mehdi 			 "
	result = strings.TrimSpace(name)
	output("trimSpace", name, result)

	name = "mehdi_teymorian"
	suffix := strings.HasSuffix(name, "mehdi")
	output("hasSuffix", name, suffix)

	name = "hello 1 world"
	result = strings.ReplaceAll(name, "1", ":)")
	output("replace", name, result)

}

func output(action, before string, after any) {
	fmt.Println("action:", action, " == ", "before:", before, " == ", "after:", after)
}
