package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		inputReader *bufio.Reader
		input       string
		err         error
	)
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err = inputReader.ReadString('\n')
	//fmt.Printf(input)
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}
