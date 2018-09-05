package main

import (
"fmt"
"io/ioutil"
)

func main() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err == nil {
		fmt.Println(string(contents))
	} else {
		fmt.Println("cannot print file contents:", err)
	}
}
