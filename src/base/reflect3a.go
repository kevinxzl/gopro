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

type Manager struct {
	User
	Title string
}





func main() {
	m := Manager{User:User{1001, "KX001", 99}, Title: "CEO"}
	t := reflect.TypeOf(m)

	fmt.Printf("%#v\n", t.FieldByIndex([] int{0,1}))

}