// perulangan
package main

import "fmt"

func main(){
	
	for counter := 1; counter < 10; counter++ {
		fmt.Println("perulangan ke", counter)
		
	}

	slice :=[]string{"niko", "angga", "yeremia", "dewa", "aja"}

		for i := 0; i < len(slice); i++ {
			fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "niko"
	person["titte"] = "programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
