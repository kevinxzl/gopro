package main

import "fmt"

func main() {
	var (
		fName, lName string
	)
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&fName, &lName)
	fmt.Printf("Hi %s %s!\n", fName, lName)

}
