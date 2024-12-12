package main

import "fmt"

func main() {
	var a = 10
	var b = 20

	var hasil = a * b
	fmt.Println(hasil)

	var c = 5.0
	var d = 10.0

	var e = c / d
	fmt.Println(e)

	var i = 10
	i += 10
	fmt.Println(i)

	i += 5
	fmt.Println(i)

	var j = 1 
	j++ // j = j + 1
	fmt.Println(j) 

	j++ // j = j + 1
	fmt.Println(j)

	j-- // j = j - 1
	fmt.Println(j)
	j-- // j = j - 1
	fmt.Println(j)
}

