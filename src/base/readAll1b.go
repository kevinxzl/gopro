package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("10 20 30")

	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes)
	str := fmt.Sprintf("%s", bytes)
	fmt.Println(str)
}
