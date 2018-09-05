package main

import "fmt"

func main() {
	fmt.Println("Test case 1")
	for i := 0; i < 3; i++ {
		fmt.Println("i=",i)
		defer func() {
			fmt.Println(i)
		}()
	}
}
