package main

import "fmt"

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type Status int

const (
	Bad Status = iota
	Normal
	Good
	Best
)

func main() {

	x := Male

	y := string(x)

	fmt.Println(x)
	fmt.Println(y)

}
