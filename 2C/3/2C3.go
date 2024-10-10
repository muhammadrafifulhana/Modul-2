package main

import "fmt"

func main() {
	var bilangan int
	fmt.Print("Bilangan: ")
	fmt.Scan(&bilangan)

	if bilangan <= 1 {
		fmt.Println("Bilangan harus lebih besar dari 1")
		return
	}

	fmt.Print("Faktor: ")
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

	var prima bool = true
	if bilangan > 2 {
		for i := 2; i < bilangan; i++ {
			if bilangan%i == 0 {
				prima = false
				break
			}
		}
	}

	fmt.Printf("Prima: %t\n", prima)
}
