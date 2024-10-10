package main

import "fmt"

func isTahunKabisat(tahun int) bool {
	if tahun%4 == 0 {
		if tahun%100 == 0 {
			if tahun%400 == 0 {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	} else {
		return false
	}
}

func main() {
	var tahun int
	fmt.Print("Masukkan tahun yang ingin diperiksa: ")
	fmt.Scanln(&tahun)
	if isTahunKabisat(tahun) {
		fmt.Printf("Tahun: %d\nKabisat: true\n", tahun)
	} else {
		fmt.Printf("Tahun: %d\nKabisat: false\n", tahun)
	}
}
