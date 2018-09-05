package main

import "fmt"

type A struct {
	B
}

type B struct {
	Name string
}

func main() {
	a := A {B: B{Name:"BBB"}}
	fmt.Println(a)
	fmt.Println(a.Name)
	fmt.Println(a.B.Name)
}

