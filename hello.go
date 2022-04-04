package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	displayIntro()
	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Displaying logs")
		case 0:
			fmt.Println("Exiting application")
			os.Exit(0)
		default:
			fmt.Println("Command not found")
			os.Exit(-1)
		}
	}

}

func displayIntro() {
	name := "John"
	age := 32
	version := 1.18
	fmt.Println("Hello, Sir", name, "you are", age, "years old")
	fmt.Println("This Program is on the version", version)
}

func showMenu() {
	fmt.Println("1 - Initiate Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit Program")
}

func readCommand() int {
	var accessCommand int
	fmt.Scan(&accessCommand)
	fmt.Println("The chosen command was", accessCommand)

	return accessCommand
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "This was a Success.", resp.StatusCode)
	} else {
		fmt.Println("Site:", site, "has something wrong. Status Code:", resp.StatusCode)
	}

}
