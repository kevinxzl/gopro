package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Test Case1")
	m := map[int]string {112 : "a", 12 : "b", 3 : "c", 456 : "d", 15 : "e"}
	s := make([]int, len(m))
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	sort.Ints(s)
	for _, key := range s {
		v := m[key]
		fmt.Println(key,":", v)
	}
	fmt.Println(m)

	fmt.Println("Test Case 2")
	m2 := map[int]string{1 : "a", 2 : "b", 3 : "c", 4 : "d", 5 : "e"}
	m3 := make(map[string]int , len(m2))
	for k, v := range m2 {
		m3[v] = k
	}
	fmt.Println("m2:", m2)
	fmt.Println("m3:", m3)
}
