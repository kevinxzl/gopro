package main

import "fmt"

func main() {
	fs := [4] func() {}

	fmt.Println("Test Case1")
	for i := 0; i < 4; i++ {
		fmt.Println("i1=", i)
		defer fmt.Println("defer i=", i)
		defer func() {
			fmt.Println("defer_closure i = ", i)
		}()

		fs[i] = func() {
			fmt.Println("closure i = ", i)
		}
	}

	fmt.Println("Test Case2 ")
	for i, f := range fs {
		fmt.Println("i2=", i)
		f()
	}
}

