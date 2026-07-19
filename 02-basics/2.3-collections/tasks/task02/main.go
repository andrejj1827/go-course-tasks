package main

import (
	"fmt"
	"slices"
)

func invertMap(m map[string]int) map[int]string {
	result := make(map[int]string)
	for key, value := range m {
		result[value] = key
	}
	return result
}

func main() {
	fruits := map[string]int{
		"яблоко":   1,
		"банан":    2,
		"апельсин": 3,
	}

	inverted := invertMap(fruits)

	keys := make([]int, 0, len(inverted))
	for k := range inverted {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	for _, k := range keys {
		fmt.Printf("%d -> %s\n", k, inverted[k])
	}
}