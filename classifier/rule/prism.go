package rule

import (
	"fmt"
)

func Prism(path string) {
	fmt.Println("reading file...")
	f := readFile(path)
	ts := createTrainingSet(f)
	f.Close()

	rules := make(map[string][]Rule)
	for _, class := range ts.class.values {
		fmt.Println("finding rules for class " + class)
		rules[class] = nil
		recs := []int{}
		for i := 0; i < len(ts.records); i++ {
			if ts.records[i][len(ts.records[i])-1] == class {
				recs = append(recs, i)
			}
		}

		for len(recs) != 0 {
			var rule Rule
			pure := false
			cov := []int{}
			for i := 0; i < len(ts.records); i++ {
				cov = append(cov, i)
			}
			attrsInfo := initAttrsInfo(ts.attributes)
			for !(pure || len(rule) == len(ts.attributes)) {
				nc := findNewCondition(cov, ts.records, attrsInfo, ts.attributes, class)
				rule = append(rule, condition{attribute: nc.attr, value: nc.val})
				if nc.ratio == 1 {
					pure = true
				}
				delete(attrsInfo, nc.attr)
				for _, values := range attrsInfo {
					for _, nums := range values {
						nums[0] = 0
						nums[1] = 0
					}
				}
				filtered := []int{}
				for _, rec := range cov {
					if ts.records[rec][ts.attrToIndex[nc.attr]] == nc.val {
						filtered = append(filtered, rec)
					}
				}
				cov = filtered
			}
			//remove covered instances by the rule
			for _, i := range cov {
				recs[i] = -1
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
