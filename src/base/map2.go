package main

import "fmt"

func main() {
	fmt.Println("Test Case 01")
	var m1 map[int]map[int]string
	m1 = make(map[int]map[int]string)
	m1[1] = make(map[int]string)
	m1[1][1] = "OK"
	a1 := m1[1][1]
	fmt.Println(a1)

	fmt.Println("Test Case 02")
	m2 := make(map[int]map[int]string)
	if _, flag := m2[2][1]; !flag {
		m2[2] = make(map[int]string)
	}
	m2[2][1] = "GOOD"
	a2, flag := m2[2][1]
	fmt.Println(a2, flag)

	fmt.Println("Test Case 03")
	sm3 := make([]map[int]string, 5)
	for i, v := range sm3 {
		v = make(map[int]string, 1)
		v[i] = "OK"
		fmt.Println(v)
	}
	fmt.Println(sm3)

	fmt.Println("Test Case 04")
	sm4 := make([]map[int]string, 5)
	for i := range sm4 {
		sm4[i] = make(map[int]string, 1)
		sm4[i][i] = "GOOD"
		fmt.Println(sm4[i])
	}
	fmt.Println(sm4)
}
