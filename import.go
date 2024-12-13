package main

import "fmt"
import "belajar-golang-dasar/helper"

func main() {
	result := helper.SayHello("Eko")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version)
	// fmt.Println(helper.sayGoodbye("Eko"))

}