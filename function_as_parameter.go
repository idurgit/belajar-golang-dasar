package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string)  {
	filteredName := filter(name) 
	fmt.Println("Hello", filteredName) 
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Eko", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("anjing", filter)
}