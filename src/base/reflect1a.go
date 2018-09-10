package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x1 float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x1))
	x2 := 99
	fmt.Println("type:", reflect.TypeOf(x2))
}
