package main

import "fmt"

func logging() {
	fmt.Println("selesai memangil function")
}
func runApplication(value int) {
	defer logging()
	fmt.Println("run Application")
	result := 10 / value
	fmt.Println("result", result)
}

func main() {
	runApplication(2)
}