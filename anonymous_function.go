package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList)  {
	if blackList(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome to our system", name)
	}
}

func main() {
	BlackList := func(name string) bool {
		return name == "anjing"
	} 

	registerUser("Eko", BlackList)
	registerUser("anjing", BlackList)
}