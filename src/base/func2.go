package main

import "fmt"

func closure(x int) func(int) int {
	fmt.Printf("x addr =%p\n", &x)
	return func(y int) int {
		fmt.Printf("a address =%p\n", &x)
		return x + y
	}
}

func main() {
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}
