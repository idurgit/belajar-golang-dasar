package main

import "fmt"

func getGoodbye(name string) string {
	return "Goodbye " + name
}

func main() {
	contoh := getGoodbye
	misal := getGoodbye

	fmt.Println(contoh("Eko"))
	fmt.Println(misal("Eko"))
}
