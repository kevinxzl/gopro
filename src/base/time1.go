package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Origin:", t)
	fmt.Println("ANSIC:", t.Format(time.ANSIC))
}