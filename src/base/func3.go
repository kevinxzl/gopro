package main

import "fmt"

func main() {

	fmt.Println("Test Case 1")
	for i := 0; i < 3; i++ {
		fmt.Println("i=",i)
		defer fmt.Println(i)
	}
}
