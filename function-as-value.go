package main

import "fmt"

func getGoodBye(name string)string {
	return "good bye " + name
}

func main() {
	sayGoodBye := getGoodBye

	result := sayGoodBye("Niko")
	fmt.Println(result)
	fmt.Println(sayGoodBye("Niko"))
}