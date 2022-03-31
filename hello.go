package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "John"
	age := 32
	version := 1.18
	fmt.Println("Hello, Sir", name, "you are", age, "years old")
	fmt.Println("This Program is on the version", version)

	fmt.Println("The type of the variable name is", reflect.TypeOf(version))
}
