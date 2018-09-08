package main

import (
	"fmt"
	"time"
)



func main() {

	s := []string {"a", "b", "c", "d"}
	for i, v := range s {
		go func(i int, v string) {
			fmt.Println("Thread: ", i, "---", v)
		}(i, v)
	}

	time.Sleep(time.Second)
}
