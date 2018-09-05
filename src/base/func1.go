package main

import "fmt"

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func sum(nums ...int) int {
	s := 0
	for i := range nums {
		s += nums[i]
	}
	return s
}


func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}
}

func main() {
	fmt.Println("Test01: 13/3 = 4...1")
	q, r := div(13, 3)
	fmt.Printf("13 div 3 is %d mod %d\n", q, r)

	fmt.Println("Test02")
	fmt.Println(sum(1,2,3,4,5,6,7,8,9))

	fmt.Println("Test Case03")
	if result, err := eval(3, 4, "*"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}


