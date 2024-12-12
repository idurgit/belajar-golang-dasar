package main

import "fmt"

func main() {
	var name [3]string
	name[0] = "Jaka"
	name[1] = "Kelana"
	name[2] = "Umbara"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [...]int{90, 80, 70, 100, 110}
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))
}