package main

import "fmt"

func A() {
	fmt.Println("Func A")
}

func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B", "for ",  err)
		}
	}()

	panic("Panic in Func B")
}

func C() {
	fmt.Println("Func C")
}

func main() {
	A()
	B()
	C()
}

