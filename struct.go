package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	var eko Customer
	fmt.Println(eko)

	eko.Name = "Eko Kurniawan"
	eko.Address = "Jakarta"
	eko.Age = 30

	fmt.Println(eko)
	fmt.Println(eko.Name) 
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)

	joko := Customer{"Joko", "Bandung", 20}
	fmt.Println(joko)

	budi := Customer{"Budi", "Semarnag", 25}
	fmt.Println(budi)

	eko.sayHello("Budi")
}