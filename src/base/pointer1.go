package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {

	fmt.Println("Test Case01")
	var a int = 2
	var pa *int = &a
	*pa = 100
	fmt.Println(a)

	fmt.Println("Test Case02")
	a2, b2 := 100, 999
	fmt.Println(a2, b2)
	swap(&a2, &b2)
	fmt.Println(a2, b2)

}
