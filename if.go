package main

import "fmt"

func main() {
	name := "Eko"

	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hi, what's your name?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
} else {
		fmt.Println("Nama sudah benar")
	}
}