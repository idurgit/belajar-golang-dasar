package main

import "fmt"

func endApp() {
	fmt.Println("End app")
	message := recover()
	fmt.Println("Error Message : ", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Ups Error")
	}
}

func main() {
	runApp(true)
	fmt.Println("End Program")
}
