package main

import "fmt"



func main() {
	c := make(chan bool, 3)
	go func() {
		fmt.Println("Go Go Go !!!")
		c <- true
		c <- false
		c <- true
		//close(c)
	}()

	//using for range to wait rotoutine, must close chan in the thread
	for i := 0; i < 3; i++ {
		u := <- c
		fmt.Println(u)
	}
}
