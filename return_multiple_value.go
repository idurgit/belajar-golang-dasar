package main

import "fmt"

func getFullName() (string, string) {
	return "Eko", "Kurniawan"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}