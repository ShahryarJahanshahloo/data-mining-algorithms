package main

import (
	// "fmt"

	"github.com/bsm/arff"
)

//TODO: we can store all the data in our own vars and close the file

// assuming that the last attribute is the class
// may replace with a map
type Attribute struct {
	name   string
	values []string
}

var attributes []Attribute

func main() {
	//reading the input file
	data, err := arff.Open("./contact-lenses.arff")
	if err != nil {
		panic("failed to open file: " + err.Error())
	}
	defer data.Close()

	//preparing class
	class := Attribute{}
	attributesLenghth := len(data.Attributes)
	class.name = data.Attributes[attributesLenghth-1].Name
	class.values = data.Attributes[attributesLenghth-1].NominalValues

	//storing other variables and their values
	for index := 0; index < attributesLenghth-1; index++ {
		attr := data.Attributes[index]
		attributes = append(attributes, Attribute{name: attr.Name, values: attr.NominalValues})
	}

	// for index, attr := range attributes {
	// 	var coveredDocs []int
	// }

	// for data.Next() {
	// 	fmt.Println(data.Row().Values...)
	// }
	if err := data.Err(); err != nil {
		panic("failed to read file: " + err.Error())
	}
}
