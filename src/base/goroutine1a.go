package main

import "fmt"



func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("Go Go Go !!!")
		c <- true
		c <- false
		c <- true
		close(c)
	}()

	//using for range to wait rotoutine, must close chan in the thread
	for v := range c {
		fmt.Println(v)
	}
}
