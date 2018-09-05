package main

import "fmt"

type person struct {
	Name string
	Age int
}

type personA struct {
	string
	int
}

func A(p person) {
	p.Age = 1000
	p.Name = "NO BODY"
}

func B(p *person) {
	p.Name = "NO BODY"
	p.Age = 10
}

func main() {

	fmt.Println("Test Case1 1")
	p1 := person{}
	fmt.Println(p1)
	p1.Name = "KX"
	p1.Age = 18
	fmt.Println(p1)

	fmt.Println("Test Case 2")
	p2 := person {
		Name: "Peter",
		Age: 99,
	}
	fmt.Println(p2)

	fmt.Println("Test Case 3")
	A(p2)
	fmt.Println(p2)

	fmt.Println("Test Case 4")
	B(&p2)
	fmt.Println(p2)

	fmt.Println("Test Case 5")
	p5 := personA{"KX05", 19}
	fmt.Println(p5)


}
