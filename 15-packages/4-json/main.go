package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrFailMarshal   = errors.New("failed to marshal")
	ErrFailIndent    = errors.New("failed to indent")
	ErrFailUnmarshal = errors.New("failed to unmarshal")
)

type Person struct {
	Name         string `json:"name"`
	Age          int    `json:"age"`
	LuckyNumbers []int  `json:"lucky_numbers"`
}

func main() {

	p := Person{
		Name:         "Mehdi",
		Age:          30,
		LuckyNumbers: []int{1, 2, 3},
	}
	bytesResult, err := json.Marshal(p)
	if err != nil {
		panic(ErrFailMarshal)
	}

	fmt.Println(string(bytesResult))

	indentedResult := new(bytes.Buffer)
	err = json.Indent(indentedResult, bytesResult, "", "\t")
	if err != nil {
		panic(ErrFailIndent)
	}

	fmt.Println(indentedResult.String())

	fmt.Println("====================")

	raw := `{"name":"Ali","age":5,"lucky_numbers":[10]}`

	ali := new(Person)

	err = json.Unmarshal([]byte(raw), ali)
	if err != nil {
		panic(ErrFailUnmarshal)
	}

	fmt.Printf("%+v\n", *ali)

}
