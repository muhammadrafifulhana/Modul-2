package main

import "fmt"

func main() {
	var r int
	fmt.Print("Jari-Jari: ")
	fmt.Scanln(&r)

	const pi float64 = 3.1415926535

	volumeBola := (4.0 / 3.0) * pi * float64(r*r*r)
	luasBola := 4.0 * pi * float64(r*r)

	fmt.Printf("Bola dengan Jari-Jari %d memiliki volume %.4f dan luas kulit %.4f\n", r, volumeBola, luasBola)
}
