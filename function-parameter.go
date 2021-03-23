package main

import "fmt"

func sayHelloTo(firstName string, lastName string){
	fmt.Println("hello", firstName, lastName);
}
func main(){
	firstName := "Niko"
	sayHelloTo(firstName, "Demus")
	sayHelloTo("Angga", "Kamiswara")
	}
