package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 10
	s[1] = 100
	s[2] = 200
}

func main() {
	arr := [...]int {0, 1, 2, 3, 4, 5, 6, 9, 8, 7}

	fmt.Println("Test Case1")
	fmt.Println("arr[2:6]=", arr[2:6])
	fmt.Println("arr[:6]=", arr[:6])
	fmt.Println("arr[2:]=", arr[2:])
	fmt.Println("arr[:]=", arr[:])

	fmt.Println("Test Case2")
	s1 := arr[:]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	updateSlice(s1)
	fmt.Println("after update slice:", arr)
	s2 := arr[3:6]
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
	fmt.Println("Test Case3")
	s3a := arr[2:6]
	s3b := s3a[3:5]
	fmt.Println("arr:", arr)
	fmt.Println("s3a:", s3a)
	fmt.Println("s3b s3a[3:5]",s3b)
}
