package main

import "fmt"

func main() {

	var nam float64
	var nmk string
	fmt.Print("Nilai akhir mata kuliah: ")

	fmt.Scanln(&nam)
	if nam > 80 {
		nmk = "A"
	}
	if nam > 72.5 {
		nmk = "AB"
	}
	if nam > 65 {
		nmk = "B"
	}
	if nam > 57.5 {
		nmk = "BC"
	}
	if nam > 50 {
		nmk = "c"
	}
	if nam > 40 {
		nmk = "D"
	} else if nam <= 40 {
		nmk = "F"
	}

	fmt.Println("Nilai mata kuliah: ", nmk)
}

/*
	a. Jika nam diberikan adalah 80.1, apa keluaran dari program tersebut? Apakah eksekusi program tersebut sesuai spesifilasi soal?
	Jika input nam adalah 80.1, maka output dari program tersebut adalah "AB". Namun, ini tidak sesuai dengan spesifikasi soal. Output yang benar seharusnya adalah "A" karena 80.1 lebih besar dari 80.

	b. Apa saja Resalahan dari program tersebut? Mengapa demikian? Jelaskan alur program seharusnya!
	Kesalahan dalam program adalah:

	Program membandingkan string nam dengan nilai numerik menggunakan operator >, yang salah. Variabel nam harus dikonversi ke tipe float64 sebelum dibandingkan.
	Perintah if-else tidak saling eksklusif, artinya jika kondisi pertama benar, maka kondisi berikutnya juga akan dievaluasi dan menggantikan nilai sebelumnya. Ini bukan perilaku yang diinginkan.

	Alur program seharusnya adalah:
	- Baca input nam sebagai tipe float64.
	- Gunakan perintah if-else dengan kondisi yang saling eksklusif untuk menentukan nilai berdasarkan nilai input.
	- Tetapkan nilai yang sesuai ke variabel nmk.
*/

// c. Perbaiki program tersebut! Ujilah dengan masukan: 93.5; 70.6; dan 49.5. Seharusya keluaran yang diperoleh adalah 'A', 'B', dan 'D'

package main

import "fmt"

func main() {
    var nam float64
    var nmk string
    fmt.Print("Nilai akhir mata kuliah: ")
    fmt.Scanln(&nam)

    if nam >= 80 {
        nmk = "A"
    } else if nam >= 72.5 {
        nmk = "B"
    } else if nam >= 65 {
        nmk = "B"
    } else if nam >= 57.5 {
        nmk = "BC"
    } else if nam >= 50 {
        nmk = "C"
    } else if nam >= 40 {
        nmk = "D"
    } else {
        nmk = "F"
    }

    fmt.Println("Nilai mata kuliah: ", nmk)
}
