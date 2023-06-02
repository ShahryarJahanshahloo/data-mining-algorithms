package rule

import (
	"github.com/bsm/arff"
)

func readFile(path string) *arff.Reader {
	f, err := arff.Open(path)
	if err != nil {
		panic(err.Error())
	}
	return f
}

type attribute struct {
	name   string
	values []string
}

type trainingSet struct {
	records     [][]any
	class       attribute
	attributes  []attribute
	dist        map[any]int
	attrToIndex map[any]int
}

func createTrainingSet(f *arff.Reader) trainingSet {
	ts := trainingSet{}
	l := len(f.Attributes)
	ts.class.name = f.Attributes[l-1].Name
	ts.class.values = f.Attributes[l-1].NominalValues
	ts.dist = make(map[any]int)
	ts.attrToIndex = make(map[any]int)
	for _, v := range ts.class.values {
		ts.dist[v] = 0
	}
	for i := 0; i < l-1; i++ {
		a := f.Attributes[i]
		ts.attrToIndex[a.Name] = i
		ts.attributes = append(ts.attributes, attribute{name: a.Name, values: a.NominalValues})
	}
	for f.Next() {
		vs := f.Row().Values
		ts.records = append(ts.records, vs)
		ts.dist[vs[len(vs)-1]] += 1
	}
	err := f.Err()
	if err != nil {
		panic(err.Error())
	}
	return ts
}

type attrValueInfo map[any]map[any][]int

// assuming the first one is valid, second one total
func initAttrsInfo(attrs []attribute) attrValueInfo {
	res := make(attrValueInfo)
	for _, attr := range attrs {
		res[attr.name] = make(map[any][]int)
		for _, value := range attr.values {
			res[attr.name][value] = []int{0, 0}
		}
	}
	return res
}

type condition struct {
	attribute any
	value     any
}

type Rule []condition

// func findNewRule() {}

type newCondition struct {
	attr    any
	val     any
	ratio   float32
	covered int
}

func findNewCondition(under []int, records [][]any, table attrValueInfo, attrs []attribute, class string) newCondition {
	for _, rec := range under {
		for i := 0; i < len(records[rec])-1; i++ {
			if len(table[attrs[i].name]) != 0 {
				table[attrs[i].name][records[rec][i]][1] += 1
				if records[rec][len(records[rec])-1] == class {
					table[attrs[i].name][records[rec][i]][0] += 1
				}
			}
		}
	}
	best := newCondition{"", "", 0, 0}
	for attr, values := range table {
		for value, numbers := range values {
			var new float32 = float32(numbers[0]) / float32(numbers[1])
			if new > best.ratio {
				best.attr = attr
				best.val = value
				best.ratio = new
				best.covered = numbers[1]
			} else if new == best.ratio {
				if numbers[1] > best.covered {
					best.attr = attr
					best.val = value
					best.ratio = new
					best.covered = numbers[1]
				}
			}
		}
	}
	return best
}
