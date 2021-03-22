package main

import "fmt"

func main(){
	var ujian = 80
	var absensi = 80

	var lulusUjian = ujian >= 80
	var luluAbsensi = absensi >= 80
	fmt.Println(lulusUjian)
	fmt.Println(luluAbsensi)

	var lulus = luluAbsensi && lulusUjian
	fmt.Println(lulus)

	fmt.Println(ujian >= 80 && absensi 
	>= 80)
}
