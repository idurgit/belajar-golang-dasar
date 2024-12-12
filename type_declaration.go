package main

import "fmt"

func main() {
	type NoKTP string
	var ktpEko NoKTP = "11111111111111"
	var contoh = "1234567890"
	var contohKtp = NoKTP(contoh)

	fmt.Println(ktpEko)
	fmt.Println(contohKtp)
	
}