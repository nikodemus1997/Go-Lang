package main

import "fmt"

func main(){
	type noKTP string
	type merried bool

	var noKTPNiko noKTP = "7340723047020374"
	var merriedStatus merried = true

	fmt.Println(noKTPNiko)
	fmt.Println(merriedStatus)
}
