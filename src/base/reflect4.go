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

func (u User)Hello(name string){
	fmt.Println("Hello", name, "My name is ", u.Name)
}


func main() {

	u := User{1001, "KX01", 99}
	u.Hello("Peter")

	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("Joe")}
	mv.Call(args)
}