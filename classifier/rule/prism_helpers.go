package rule

import (
	"fmt"
	"github.com/bsm/arff"
)

type Attribute struct {
	name   string
	values []string
}

type DataSet struct {
	records    [][]any
	class      Attribute
	attributes []Attribute
	pairs      int
	dist       map[any]int
}

func readFile(path string) (dataset DataSet, err error) {
	data, err := arff.Open(path)
	if err != nil {
		return
	}
	defer data.Close()

	attributesLenghth := len(data.Attributes)
	dataset.class.name = data.Attributes[attributesLenghth-1].Name
	dataset.class.values = data.Attributes[attributesLenghth-1].NominalValues
	dataset.dist = make(map[any]int)
	for _, v := range dataset.class.values {
		dataset.dist[v] = 0
	}
	for index := 0; index < attributesLenghth-1; index++ {
		attr := data.Attributes[index]
		dataset.pairs += len(attr.NominalValues)
		dataset.attributes = append(dataset.attributes, Attribute{name: attr.Name, values: attr.NominalValues})
	}

	for data.Next() {
		values := data.Row().Values
		dataset.records = append(dataset.records, values)
		dataset.dist[values[len(values)-1]] += 1
	}
	err = data.Err()
	if err != nil {
		return
	}
	return
}

// assuming the first one is valid, second one total
func initAttrInfo(attrs []Attribute) map[string]map[string][2]int {
	res := make(map[string]map[string][2]int)
	for _, attr := range attrs {
		res[attr.name] = nil
		for _, value := range attr.values {
			res[attr.name][value] = [2]int{0, 0}
		}
	}
	return res
}

type Condition struct {
	attribute string
	value     string
}

type Rule []Condition

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

func findNewRule() {}

func findNewCondition(records [][]any) {
	for _, rec := range records {
		for index := 0; index < len(rec)-1; index++ {
			fmt.Println("break")
			break
		}
	}
}
