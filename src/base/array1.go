package main

import "fmt"

func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {

	fmt.Println("Test Case1")
	var arr11 [5]int
	arr12 := [3]int {1, 3, 5}
	arr13 := [...] int {2, 4, 6, 8, 10}
	var arr14 [5][3] int

	fmt.Println(arr11, arr12, arr13)
	fmt.Println(arr14)

	fmt.Println("Test Case2")
	for i2 := range arr13 {
		fmt.Println(arr13[i2])
	}

	fmt.Println("Test Case3")
	for i3, v3 := range arr13 {
		fmt.Println(i3, v3)
	}

	fmt.Println("Test Case4")
	for _, v4 := range arr13 {
		fmt.Println(v4)
	}

	fmt.Println("Test case5")
	printArray(arr11)
	fmt.Println("Test case6")
	printArray(arr13)
}
