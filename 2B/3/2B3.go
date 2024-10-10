package main

import "fmt"

func main() {
	for {
		var beratKiri, beratKanan float64
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&beratKiri, &beratKanan)

		if beratKiri < 0 || beratKanan < 0 || beratKiri+beratKanan > 150 {
			break
		}

		if abs(beratKiri-beratKanan) >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}
	fmt.Println("Proses selesai.")
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
