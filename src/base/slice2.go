package main

import "fmt"

func main() {
	arr := [...]int {0, 11, 22, 33, 44, 55, 66, 77}
	s0 := arr[:]
	s0 = append(s0, 9)
	s1 := arr[2:6]
	s2 := s1[3:5]
	s3 := append(s2, 200)
	s4 := append(s3, 300)
	s5 := append(s4, 400)


	fmt.Println("arr=", arr)
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)
	fmt.Println("s3=", s3)
	fmt.Println("s4=", s4)
	fmt.Println("s5=", s5)
	fmt.Println("s0=", s0)
}
