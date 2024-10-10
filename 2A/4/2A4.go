package main

import "fmt"

func main() {
	var celcius float32

	fmt.Print("Temperatur Celcius : ")
	fmt.Scanln(&celcius)

	fahrenheit := (celcius * 9 / 5) + 32

	reamur := celcius * 4 / 5

	kelvin := (fahrenheit + 459.67) * 5 / 9

	hasilFahrenheit := fahrenheit
	hasilReamur := reamur
	hasilKelvin := kelvin

	fmt.Printf("Derajat Reamur : %d \n", int(hasilReamur))
	fmt.Printf("Derajat Fahrenheit : %d \n", int(hasilFahrenheit))
	fmt.Printf("Derajat Kelvin : %d \n", int(hasilKelvin))
}
