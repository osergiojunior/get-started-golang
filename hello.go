package main

import (
	"fmt"
)

func main() {
	name := "John"
	age := 32
	version := 1.18
	fmt.Println("Hello, Sir", name, "you are", age, "years old")
	fmt.Println("This Program is on the version", version)

	fmt.Println("1 - Initiate Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit Program")

	var command int
	fmt.Scan(&command)
	fmt.Scanf("%d", &command)
	fmt.Println("The chosen command was", command)

}
