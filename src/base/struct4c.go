package main

import "fmt"

type A struct {
	B
	C
}

type B struct {
	Name string
}

type C struct {
	Name string
}

func main() {
	a := A {B: B{Name:"BBB"}, C:C{Name:"CCC"}}
	fmt.Println(a)
	fmt.Println(a.C.Name, a.B.Name)

}

