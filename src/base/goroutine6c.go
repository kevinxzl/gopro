package main

import (
	"fmt"
	"sync"
)



func main() {

	s := []string {"a", "b", "c", "d"}
	wg := sync.WaitGroup{}
	wg.Add(4)
	for i, v := range s {
		go func(wg *sync.WaitGroup, i int, v string) {
			fmt.Println("Thread: ", i, "---", v)
			wg.Done()
		}(&wg, i, v)
	}

	wg.Wait()
}
