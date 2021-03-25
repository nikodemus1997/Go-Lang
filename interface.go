package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}
func main() {
	var niko Person
	niko.Name = "Niko"

	sayHello(niko)

	cat := Animal{
		Name: "Push",
	}
	sayHello(cat)
}