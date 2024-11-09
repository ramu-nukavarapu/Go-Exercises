package main

import "fmt"

func main() {
	var fahrenheitValue float64
	fmt.Print("Enter fahrenheit value: ")
	fmt.Scanf("%f", &fahrenheitValue)

	fmt.Println("Celsius value is", FahrenheitToCelsius(fahrenheitValue))
}

func FahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}
