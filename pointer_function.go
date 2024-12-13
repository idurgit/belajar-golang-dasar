package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"

	fmt.Println(address)
}

func main() {
	address := new(Address)

	ChangeCountryToIndonesia(address)

	fmt.Println(address)
}