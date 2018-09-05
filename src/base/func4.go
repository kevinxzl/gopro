package main

import "fmt"

func A() {
	fmt.Println("Func A")
}

func B() {
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