package main

import (
	"fmt"
	"reflect"
)

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

func main() {
	var n int
	fmt.Println("Please input a integer")
	fmt.Scanf("%d", &n)
	fmt.Println(n)
//	fmt.Println("type:", reflect.TypeOf(n))

	f := fibonacci()
	for i := 0; i <= n; i++ {
		fmt.Println(f())
	}


}
