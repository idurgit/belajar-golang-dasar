package main

import "fmt"

func main() {
	nama := "Kurniawan" 
	switch nama {
	case "Eko":
		fmt.Println("Hello Eko")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hi, what's your name?")
	}

	switch length := len(nama); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

	length := len(nama)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}