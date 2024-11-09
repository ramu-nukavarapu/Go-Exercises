package main

import "fmt"

func main() {
	var value int
	fmt.Print("Enter a value: ")
	fmt.Scanf("%d", &value)

	if value%2 == 0 {
		fmt.Printf("%v is even\n", value)
	} else {
		fmt.Printf("%v is odd\n", value)
	}
}
