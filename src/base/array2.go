package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [] int {10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	s := arr
	sort.Ints(s)
	fmt.Println(s)

	for i, v := range arr {
		v = i * 2 + v * 3
		fmt.Println(v)
	}
	fmt.Println(arr)
}
