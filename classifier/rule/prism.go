package rule

import (
	"fmt"
)

func Prism(path string) {
	dataset, err := readFile(path)
	if err != nil {
		panic(err.Error())
	}

	rules := make(map[string][]Rule)
	for _, classValue := range dataset.class.values {
		rules[classValue] = nil
		covered := []int{}
		for len(covered) != dataset.dist[classValue] {
			var newRule Rule
			isPure := false
			underConditions := make([][]any, len(dataset.records))
			copy(underConditions, dataset.records)
			// attributesInfoGain := initAttrInfo(dataset.attributes)
			for !(isPure || len(newRule) == len(dataset.attributes)) {
				findNewCondition(underConditions)
			}
		}
	}

	//printing res
	fmt.Println(dataset)
	for class, rules := range rules {
		fmt.Println("Rules found for class = " + class)
		for _, r := range rules {
			fmt.Println(r)
		}
	}
}
