package main

import "fmt"

func main() {
	m1 := map[string]string {
		"name": "KX",
		"course": "golang",
		"site": "kx.com",
		"email": "kx01@qq.com",
		"xx": "ABC",
	}

	delete(m1, "xx")

	m2 := make(map[string]int)  //m2 == empty map

	var m3 map[string]int  // me = nil

	fmt.Println(m1, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	fmt.Println("Traversing map key ")
	for k1 := range m1 {
		fmt.Println(k1)
	}

	fmt.Println("Traversing map value")
	for _, v2 := range m1 {
		fmt.Println(v2)
	}

}

