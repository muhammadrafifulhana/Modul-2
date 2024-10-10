package main

import (
	"fmt"
	"math"
)

func main() {
	var k int
	fmt.Print("Masukkan nilai K: ")
	fmt.Scanln(&k)

	akar2 := hitungAkar2(k)
	fmt.Printf("Nilai akar 2 untuk K = %d adalah %.10f\n", k, akar2)
}

func hitungAkar2(k int) float64 {
	var hasil float64 = 1.0
	for i := 0; i < k; i++ {
		hasil *= math.Pow((4*float64(i)+2), 2) / ((4*float64(i) + 1) * (4*float64(i) + 3))
	}
	return hasil
}
