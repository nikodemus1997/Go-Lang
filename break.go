package main

import "fmt"

func main(){
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // untuk menghentikan perulangan
		}
		fmt.Println("perulangan ke", i );
	}
}
