package main

import (
	"fmt"
	"sort"
)

func main() {
	m := make(map[string]interface{})
	m = map[string]interface{}{
		"alice":   31,
		"charlie": 34,
		"bob":     "ssff",
		"david":   89,
	}
	sortMap(m)

}

func sortMap(m map[string]interface{}) {
	name := []string{}
	for n := range m {
		name = append(name, n)
	}
	sort.Strings(name)
	for _, v := range name {
		fmt.Printf("%s : %v \n", v, m[v])
	}
}
