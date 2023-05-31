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
	records     [][]any
	class       Attribute
	attributes  []Attribute
	pairs       int
	dist        map[any]int
	attrToIndex map[any]int
}

func readFile(path string) (dataset DataSet, err error) {
	fmt.Println("reading file...")
	data, err := arff.Open(path)
	if err != nil {
		return
	}
	defer data.Close()

	attributesLenghth := len(data.Attributes)
	dataset.class.name = data.Attributes[attributesLenghth-1].Name
	dataset.class.values = data.Attributes[attributesLenghth-1].NominalValues
	dataset.dist = make(map[any]int)
	dataset.attrToIndex = make(map[any]int)
	for _, v := range dataset.class.values {
		dataset.dist[v] = 0
	}
	for index := 0; index < attributesLenghth-1; index++ {
		attr := data.Attributes[index]
		dataset.attrToIndex[attr.Name] = index
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

type attrValueInfo map[any]map[any][]int

// assuming the first one is valid, second one total
func initAttrInfo(attrs []Attribute) attrValueInfo {
	res := make(attrValueInfo)
	for _, attr := range attrs {
		res[attr.name] = make(map[any][]int)
		for _, value := range attr.values {
			res[attr.name][value] = []int{0, 0}
		}
	}
	return res
}

type Condition struct {
	attribute any
	value     any
}

type Rule []Condition

// func findNewRule() {}

type newCondition struct {
	attr    any
	val     any
	ratio   float32
	covered int
}

func findNewCondition(under []int, records [][]any, table attrValueInfo, attrs []Attribute, class string) newCondition {
	for _, rec := range under {
		for index := 0; index < len(records[rec])-1; index++ {
			if len(table[attrs[index].name]) != 0 {
				table[attrs[index].name][records[rec][index]][1] += 1
				if records[rec][len(records[rec])-1] == class {
					table[attrs[index].name][records[rec][index]][0] += 1
				}
			}
		}
	}
	bestRatio := newCondition{"", "", 0, 0}
	for attr, values := range table {
		for value, numbers := range values {
			var newRatio float32 = float32(numbers[0]) / float32(numbers[1])
			if newRatio > bestRatio.ratio {
				bestRatio.attr = attr
				bestRatio.val = value
				bestRatio.ratio = newRatio
				bestRatio.covered = numbers[1]
			} else if newRatio == bestRatio.ratio {
				if numbers[1] > bestRatio.covered {
					bestRatio.attr = attr
					bestRatio.val = value
					bestRatio.ratio = newRatio
					bestRatio.covered = numbers[1]
				}
			}
		}
	}
	return bestRatio
}
