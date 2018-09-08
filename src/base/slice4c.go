package main

import "fmt"

func MyFunc4(s [] int) {
	s[0] = 10
	s = append(s, 99)

	fmt.Println("in MyFunc4 s=",s)
}

func main() {
	s := make([]int, 2)
	fmt.Println(s)
	MyFunc4(s)
	fmt.Println(s)
}
