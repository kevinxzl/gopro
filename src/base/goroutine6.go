package main

import (
	"fmt"
	"time"
)


func main() {

	s := []string {"a", "b", "c", "d"}
	for i, v := range s {
		go func() {
			fmt.Println("Thread: ", i, "---", v)
		}()
	}

	time.Sleep(time.Second)
}
