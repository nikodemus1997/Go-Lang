package main

import "fmt"

func main(){
	person := map[string]string{
		"nama" : "Niko",
		"address" : "Lulang",
	}
	person["title"] = "programmer"

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar go-lang"
	book["author"] = "Eko"
	book["ups"] = "salah"
	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)

}
