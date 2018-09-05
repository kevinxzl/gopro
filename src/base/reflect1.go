package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)


func mypow(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function: %s with args: "+
		"(%d, %d)\n", opName, a, b)

	return op(a, b)
}

func main() {

	fmt.Println("Test Case01")
	fmt.Println("pow(3, 4) is:", mypow(
		func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

}

