package main

import (
	"fmt"
)

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x + y
		return x
	}
}

func main() {
		f := fibonacci()
		num := 9
		for i := 0; i < num; i++ {
			res := f()
			if i==8 {
				fmt.Println(res)
			}
			//fmt.Println(f())
		}

}
