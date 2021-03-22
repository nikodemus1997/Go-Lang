package main

import "fmt"

func main(){
	name := "niko"
	switch name {
	case "niko":
		fmt.Println("hello niko")
		fmt.Println("hello niko")
	case "joko":
		fmt.Println("hello joko")
		fmt.Println("hello joko")
	default:
		fmt.Println("kenalan donk")
		fmt.Println("kenalan donk")

	}

	length := len(name)
	switch {
	case length > 10 :
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama cukup panjang")
	default :
	fmt.Println("nama sudah benar")
		
	}
	

}
