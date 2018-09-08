package main

import (
	"fmt"
)



func main() {

	s := []string {"a", "b", "c", "d"}
	c := make(chan bool, 4)
	for i, v := range s {
		go func(c chan bool, i int, v string) {
			fmt.Println("Thread: ", i, "---", v)
			c <- true
		}(c, i, v)
	}

	for i := 0; i < 4; i++ {
		<- c
	}
}
