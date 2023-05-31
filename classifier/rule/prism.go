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
	for _, class := range dataset.class.values {
		fmt.Println("finding rules for class " + class)
		rules[class] = nil
		instancesWithThisClass := []int{}
		for i := 0; i < len(dataset.records); i++ {
			if dataset.records[i][len(dataset.records[i])-1] == class {
				instancesWithThisClass = append(instancesWithThisClass, i)
			}
		}

		for len(instancesWithThisClass) != 0 {
			var newRule Rule
			isPure := false
			coveredByRule := []int{}
			for i := 0; i < len(dataset.records); i++ {
				coveredByRule = append(coveredByRule, i)
			}
			attributesInfoGain := initAttrInfo(dataset.attributes)
			for !(isPure || len(newRule) == len(dataset.attributes)) {
				newCondition := findNewCondition(coveredByRule, dataset.records, attributesInfoGain, dataset.attributes, class)
				newRule = append(newRule, Condition{attribute: newCondition.attr, value: newCondition.val})
				if newCondition.ratio == 1 {
					isPure = true
				}
				delete(attributesInfoGain, newCondition.attr)
				for _, values := range attributesInfoGain {
					for _, nums := range values {
						nums[0] = 0
						nums[1] = 0
					}
				}
				filteredUnderConditions := []int{}
				for _, rec := range coveredByRule {
					if dataset.records[rec][dataset.attrToIndex[newCondition.attr]] == newCondition.val {
						filteredUnderConditions = append(filteredUnderConditions, rec)
					}
				}
				coveredByRule = filteredUnderConditions
			}
			//remove covered instances by the rule
			for _, i := range coveredByRule {
				instancesWithThisClass[i] = -1
			}
			//add new rule to rules
		}
	}

	//printing res
	// for class, rules := range rules {
	// 	fmt.Println("Rules found for class = " + class)
	// 	for _, r := range rules {
	// 		fmt.Println(r)
	// 	}
	// }
}
