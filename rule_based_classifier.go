package main

import (
	"fmt"

	"github.com/bsm/arff"
)

func main() {
	data, err := arff.Open("./contact-lenses.arff")
	if err != nil {
		panic("failed to open file: " + err.Error())
	}
	defer data.Close()

	for data.Next() {
		fmt.Println(data.Row().Values...)
	}
	if err := data.Err(); err != nil {
		panic("failed to read file: " + err.Error())
	}
}
