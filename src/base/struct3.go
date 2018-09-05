package main

import "fmt"

type human struct {
	Sex string
}

type teacher struct {
	human
	Name string
	Age int
}

type student struct {
	human
	Name string
	Age int
}

func main() {
	t1 := teacher{Name: "James", Age: 60, human: human{Sex:"M"}}
	s1 := student{Name:"KX01", Age: 40, human: human{Sex:"F"}}
	t1.Sex = "XXX"
	fmt.Println(t1)
	fmt.Println(s1)

}

