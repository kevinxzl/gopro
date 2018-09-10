package main

import "fmt"

func main() {
	var (
		str    string
		i      int
		f      float32
		input  = "56.12 / 5212 / GO Go go"
		format = "%f / %d / %s"
	)

	fmt.Sscanf(input, format, &f, &i, &str)
	fmt.Println("From the string we read: ", f, i, str)

}
