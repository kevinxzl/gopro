package main

import "fmt"

// chan has buffer---asyn
//no buffer --- syn
func main() {
	c := make(chan bool, 1)
	go func() {
		fmt.Println("Go Go Go !!!")
		<- c
	}()

	c <- true

}
