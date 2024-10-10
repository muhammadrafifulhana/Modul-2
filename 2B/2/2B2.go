package main

import "fmt"

func main() {
	var n int
	fmt.Print("N: ")
	fmt.Scanln(&n)

	pita := ""
	for i := 1; i <= n; i++ {
		var bunga string
		fmt.Printf("Bunga %d: ", i)
		fmt.Scanln(&bunga)
		if bunga == "SELESAI" {
			break
		}
		pita += bunga + " "
	}
	fmt.Println("Pita:", pita)

	jumlahBunga := 0
	for _, bunga := range pita {
		if bunga == ' ' {
			jumlahBunga++
		}
	}
	fmt.Println("Bunga:", jumlahBunga)
}
