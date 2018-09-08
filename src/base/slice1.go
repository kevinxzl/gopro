package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 10
	s[1] = 100
	s[2] = 200
	s = append(s, 8899)
	fmt.Println("IN updateSlice", s)
}

func main() {
	arr := [...]int {0, 1, 2, 3, 4, 5, 6, 9, 8, 7}

	fmt.Println("Test Case1")
	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[:6]=", arr[:6])
	fmt.Println("arr[2:]=", arr[2:])
	fmt.Println("arr[:]=", arr[:])


	fmt.Println("Test Case2")
	s2a := arr[:]
	fmt.Printf("s2a=%v, len(s2a)=%d, cap(s2a)=%d\n", s2a, len(s2a), cap(s2a))
	updateSlice(s2a)
	fmt.Println("after update slice:", arr)
	s2b := arr[3:6]
	fmt.Printf("s2b=arr[3:6]=%v, len(s2b)=%d, cap(s2b)=%d\n", s2b, len(s2b), cap(s2b))

	fmt.Println("Test Case3")
	s3a := arr[2:6]
	s3b := s3a[3:5]
	fmt.Println("arr:", arr)
	fmt.Println("s3a:", s3a)
	fmt.Println("s3b s3a[3:5]",s3b)
}
