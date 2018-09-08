package main

import "fmt"

func Go1() {
	sum := 0
	for i := 0; i < 100000000; i++ {
		sum += i
	}
	fmt.Println(sum)
}
func main() {

	for i := 0; i < 10; i++ {
		go Go1()
	}

}
