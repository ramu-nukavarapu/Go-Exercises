package main

import "fmt"

func main() {
	var val1, val2 float64
	var operator string
	fmt.Print("Enter the two values: ")
	fmt.Scanf("%f%f", &val1, &val2)
	fmt.Print("Enter operator: ")
	fmt.Scanf("%s", &operator)

	fmt.Println("The result is", calculate(val1, val2, operator))

}

func calculate(val1, val2 float64, op string) float64 {
	switch {
	case op == "+":
		return val1 + val2
	case op == "-":
		return val1 - val2
	case op == "*":
		return val1 * val2
	case op == "/":
		return val1 / val2
	default:
		return 0
	}
}
