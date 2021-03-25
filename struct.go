package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHi(name string){
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func (a Customer) sayHuu(){
	fmt.Println("huuuu a from", a.Name)
}

func main() {
	var niko Customer
	niko.Name = "Nikodemus"
	niko.Address = "Lulang"
	niko.Age = 30

	niko.sayHi("angga")
	niko.sayHuu()

	// fmt.Println(niko.Name)
	// fmt.Println(niko.Address)
	// fmt.Println(niko.Age)

	// angga := Custumer{
	// 	Name : "Angga",
	// 	Address : "Semarang",
	// 	Age : 22,
	// }
	// fmt.Println(angga)
}



