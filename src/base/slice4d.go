package main

import "fmt"

func MyFunc4(s [] int) []int {
	s = append(s, 90)
	fmt.Println("in MyFunc4",s)
	return s
}

func main() {
	s := make([]int, 0)
	fmt.Println(s)
	s = MyFunc4(s)
	fmt.Println(s)
}
