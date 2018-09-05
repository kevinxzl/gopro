package main

import "fmt"

func printLenCap(s [] int) {
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
}

func main(){
	var arr [] int
	for i:=0; i<100; i++ {
		printLenCap(arr)
		arr = append(arr, 2 * i + 1)
	}
	fmt.Println(arr)

	fmt.Println("Test Case 2")
	s1 := []int {2, 4, 6, 8, 10}
	s2 := make([]int, 16)
	s3 := make([]int, 10, 128)
	printLenCap(s1)
	printLenCap(s2)
	printLenCap(s3)
}
