package main

import "fmt"

type TZ int

func (tz *TZ)Print() {
	fmt.Println("TTTTTTTZZZZZZZ")
}
func main()  {
	//a := TZ
	var a TZ
	a.Print()

	(*TZ).Print(&a)
}
