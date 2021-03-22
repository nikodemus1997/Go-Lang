package main

import "fmt"

func main(){

	//array manual
	var name [3] string
	name[0] = "niko"
	name[1] = "angga"
	name[2] = "Dian"

	fmt.Println(name[0]);
	fmt.Println(name[1]);
	fmt.Println(name[2]);

	//array langsung 

	var value = [3]int{
		67,
		89,
		46,
	}
	fmt.Println(value)

}
