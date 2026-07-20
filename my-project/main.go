package main

import ("fmt"
	   "slices"
	   "maps"
)
func main() {

	nums := []int{5, 2, 8, 1, 9}
	m := map[string]int{"б": 2, "а": 1, "в": 3}
	keys := slices.Collect(maps.Keys(m))
	fmt.Println(keys)
	fmt.Println(nums)
	
}
