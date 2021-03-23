// mengembalikan data
package main

import "fmt"

func getHello(name string)string{
	if name == "" {
		return "Hello Bro"
	} else {
		return "hello" + name
	}
}
func main(){
	result := getHello(" Niko")
	fmt.Println(result)
	fmt.Println(getHello(""))
}
