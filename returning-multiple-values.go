// bisa atau dapat mengembalikan data lebih dari satu data
package main

import "fmt"

func getFullName() (string, string) {
	return "Niko", "Demus"
}
func main(){
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
	}

