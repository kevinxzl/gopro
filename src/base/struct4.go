package main

import "fmt"

type A struct {
	B
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{Name:"AAA", B:B{Name:"BBBB"}}
	fmt.Println(a)
	fmt.Println(a.Name)
	fmt.Println(a.B.Name)
}
