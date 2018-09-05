package main

import (
	"fmt"
	"math"
	"strconv"
)

func convertIntToBin(n int) string {
	res := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		res += strconv.Itoa(lsb) //Atoi (string to int) and Itoa (int to string).
	}
	return res
}

func swap(a, b int) (int, int) {
	return b, a
}

func findMaxArray(arr []int) (int, int) {
	maxi := -1
	maxV := math.MinInt32
	for i, v := range arr {
		if v > maxV {
			maxi, maxV = i, v
		}
	}
	return maxi, maxV
}

func main() {
	fmt.Println("Test01: Convert int (9) to binary")
	fmt.Println(convertIntToBin(9))

	fmt.Println("Test02: swap(3,9)")
	a3, b3 := 3, 9
	fmt.Println(a3, b3)
	a3, b3 = swap(a3, b3)
	fmt.Println(a3, b3)

	fmt.Println("Test03 find the max value and position in arr3")
	arr3 := [...] int {-99, 100, 200, -100, -200, 999, 600, 700, 800, 8}
	i, v := findMaxArray(arr3[:])
	fmt.Println(i, v)

}

