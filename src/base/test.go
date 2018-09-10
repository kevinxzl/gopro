package main

import "os"
import "io/ioutil"
import "fmt"

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x + y
		return x
	}
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err == nil{
		fmt.Printf(string(bytes));
		num := int(bytes)
	}
}