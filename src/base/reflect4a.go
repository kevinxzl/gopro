package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}
func Set(o interface{}){
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("Cannot Set Value")
		return
	}else {
		v = v.Elem()

	}

	if f := v.FieldByName("Name"); f.Kind() == reflect.String {
		f.SetString("MBDBK")
	}
}

func main() {
	fmt.Println("Test Case 1")
	x := 189
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(997)
	fmt.Println(x)

	fmt.Println("Test Case 2")
	u := User{1001, "KX01", 99}
	Set(&u)
	fmt.Println(u)

}