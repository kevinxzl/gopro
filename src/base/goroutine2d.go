package main

import (
	"fmt"
	"runtime"
	"sync"
)

func Go1(wg *sync.WaitGroup, index int) {
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum += i
	}
	fmt.Println("i=", index, "---", sum)
	wg.Done()

}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//fmt.Println(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Go1(&wg, i)
	}

	wg.Wait()

}
