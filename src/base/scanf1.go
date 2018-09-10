package main

import "fmt"

func main() {
	var (
		fName, lName string
		a, b int
	)
	fmt.Println("Please enter your full name: ")
	fmt.Scanf("%s %s", &fName, &lName)
	fmt.Printf("Hi %s %s!\n", fName, lName)
	fmt.Println("Please input two int")
	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("%d + %d = %d", a, b, a+b)

}
