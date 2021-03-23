package main

import "fmt"

func main() {
	name := "Niko"
	counter := 0
	increment := func() {
		name = "Budi"
		fmt.Println("increment")
		counter++
	}

	increment()
	increment()


	fmt.Println(counter)
	fmt.Println(name)
}