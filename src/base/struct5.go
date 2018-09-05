package main

import "fmt"

type A struct {
	Name string
}

type B struct {
	Name string
}

func (a A)Print() {
	a.Name = "AAA"
	fmt.Println("Class A---Print")
}
/*
func (a A)Print(int ) {

}
*/

func (b *B)Print() {
	b.Name = "BBBBB"
	fmt.Println("Calss B---Print")
}

func main() {
	a := A{}
	b := B{}
	a.Print()
	b.Print()

	fmt.Println(a)
	fmt.Println(b)
}
