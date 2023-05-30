package rule

import (
	"fmt"

	"github.com/bsm/arff"
)

//TODO: add numeric value support for rules
//NOTE: the last attr is assumed to be the class

type Attribute struct {
	name   string
	values []string
}

type Rule []struct {
	attribute string
	value     string
}

func (r Rule) String() string {
	if len(r) < 1 {
		return "nil"
	}
	res := "(" + r[0].attribute + " = " + r[0].value + ")"
	for i := 1; i < len(r); i++ {
		res += " ^ (" + r[i].attribute + " = " + r[i].value + ")"
	}
	return res
}

func Prism(path string) {
	dataset, class, attributes, err := readFile(path)
	if err != nil {
		panic(err.Error())
	}

	rules := make(map[string][]Rule)
	for _, classValue := range class.values {
		rules[classValue] = nil
		// keep track of purity
		// keep track of used attr-value pairs mentioned in rule
		// while loop based on purity or used attr-value pairs
	}

	fmt.Println(dataset, attributes)
}

func readFile(path string) (dataset [][]any, class Attribute, attributes []Attribute, err error) {
	data, err := arff.Open(path)
	if err != nil {
		return
	}
	defer data.Close()

	attributesLenghth := len(data.Attributes)
	class.name = data.Attributes[attributesLenghth-1].Name
	class.values = data.Attributes[attributesLenghth-1].NominalValues
	for index := 0; index < attributesLenghth-1; index++ {
		attr := data.Attributes[index]
		attributes = append(attributes, Attribute{name: attr.Name, values: attr.NominalValues})
	}

	for data.Next() {
		dataset = append(dataset, data.Row().Values)
	}
	err = data.Err()
	if err != nil {
		return
	}

	return
}
