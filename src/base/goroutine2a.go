package main

import (
	"fmt"
	"runtime"
)

func Go1(c chan bool, index int) {
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum += i
	}
	fmt.Println("i=", index, "---", sum)
	if index == 9 {
		c <- true
	}
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.NumCPU())
	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go Go1(c, i)
	}

	<- c

}
