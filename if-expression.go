// kata kunci untuk percabangan
package main

import "fmt"

func main(){
	var name = "Ekko"

	if name == "Ekko" {
	fmt.Println("hello ekko");
	} else if name == "joko" {
		fmt.Println("hello joko")
	}else if name == "joko" {
		fmt.Println("hello joko")
	} else {
		fmt.Println("Kenalan dink")
	} 

	// if short statement
	var length = len(name)
	if length > 5 {
		fmt.Println("terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	} 
}
