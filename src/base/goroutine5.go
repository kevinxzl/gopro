package main

import "fmt"

//var c chan string

func MyThread(c chan string) {
	i := 0
	for {
		fmt.Println("Thread, received:", <- c)
		c <- fmt.Sprintf("Thread, Send #%d", i)
		i++
	}
}

func main() {
	c := make(chan string)
	go MyThread(c)
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("Main, send #%d", i)
		fmt.Println("Main, Received: ", <-c)
	}
}
