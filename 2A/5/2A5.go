package main

import (
	"fmt"
)

func main() {
	var integers [5]int
	var chars [3]rune

	// Membaca 5 buah data integer
	for i := 0; i < 5; i++ {
		fmt.Scan(&integers[i])
	}

	// Membaca 3 buah data karakter
	for i := 0; i < 3; i++ {
		fmt.Scanf("%c", &chars[i])
	}

	// Mencetak karakter yang sesuai dengan data integer
	for i := 0; i < 5; i++ {
		fmt.Printf("%c", integers[i])
	}
	fmt.Println()

	// Mencetak 3 buah karakter setelah karakter yang dimasukkan
	for i := 0; i < 3; i++ {
		fmt.Printf("%c", chars[i]+3)
	}
	fmt.Println()
}
