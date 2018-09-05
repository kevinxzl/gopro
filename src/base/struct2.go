package main

import "fmt"

type person struct {
	Name  string
	Age    int
	Contact struct{
		Phone, City string
	}
}


func main() {
	fmt.Println("Test Case 1")
	p1 := struct {
		Name string
		Age int
	}{
		Name: "KKX",
		Age: 20,
	}

	fmt.Println(p1)

	fmt.Println("Test Case 2")
	p2 := person{
		Name: "KX02",
		Age: 99,
	}
	p2.Contact.Phone = "13899999999"
	p2.Contact.City = "PEKING"
	fmt.Println(p2)

}
