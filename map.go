package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Eko", 
		"address": "Jakarta",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Eko"
	book["ups"] = "salah"
	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)
}
