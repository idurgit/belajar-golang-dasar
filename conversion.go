package main 

import "fmt"

func main() {
	var nilai32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32, nilai64, nilai16)

	var name = "Jaka Kelana"
	var e byte = name[0]
	var eString string = string(e)

	fmt.Println(name, e, eString)
}