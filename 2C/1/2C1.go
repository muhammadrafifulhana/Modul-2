package main

import "fmt"

func main() {
	var berat int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	var kg int = berat / 1000
	var sisaBerat int = berat % 1000

	var biayaTambahan int
	if sisaBerat >= 500 {
		biayaTambahan = sisaBerat * 5
	} else if sisaBerat > 0 {
		biayaTambahan = sisaBerat * 15
	} else {
		biayaTambahan = 0
	}

	if kg > 10 {
		biayaTambahan = sisaBerat * 5
	}

	var totalBiaya int = kg*10000 + biayaTambahan

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisaBerat)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", (kg * 10000), biayaTambahan)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
