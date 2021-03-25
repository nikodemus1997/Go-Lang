package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("terjadi error", message)
	}
	fmt.Println("aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("aplikasi berjalah")
}

func main() {
	runApp(false)
}
