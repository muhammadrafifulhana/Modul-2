package main

import "fmt"

func main() {

	/*
		variabel satu, dua, dan tiga dideklarasikan sebagai string, yang akan menyimpan input pengguna
		temp adalah variabel string sementara yang digunakan untuk menukar nilai-nilai
	*/
	var (
		satu, dua, tiga string
		temp            string
	)

	/*
		Fungsi fmt.Scanln membaca input user dan menyimpannya dalam variabel yang sesuai (satu, dua, dan tiga)
	*/
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)

	/* menampilkan nilau output asli */
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)

	/*
		menukar nilai setiap variabel secara terbalik
	*/
	temp = satu // temp mengambil nilai satu
	satu = dua  // satu mengambil nilai dua
	dua = tiga  // dua mengambil nilai tiga
	tiga = temp // tiga mengambil nilai temp (yang merupakan nilai asli satu)
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
