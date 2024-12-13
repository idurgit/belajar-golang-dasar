package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string  { 
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string  { 
	return animal.Name
}

func sayHello(value HasName)  { 
	fmt.Println("Hello", value.GetName())
}

func main() {
	person := Person{Name: "Eko Kurniawan"}
	sayHello(person) 

	animal := Animal{Name: "Kucing"}
	sayHello(animal)
}