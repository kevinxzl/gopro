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
	c <- true

}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//fmt.Println(runtime.NumCPU())
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Go1(c, i)
	}

	for i := 0; i < 10; i++ {
		<-c
	}

}
