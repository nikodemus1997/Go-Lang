package main

import "fmt"

func main(){
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue // menghentikan perulangan saat ini dan akan langsung dilanjutkan  false statement
		}
		fmt.Println("perulangan ke", i );
	}
}
