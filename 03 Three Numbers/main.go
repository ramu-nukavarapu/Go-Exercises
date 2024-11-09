package main

import "fmt"

func main() {
	var val1, val2, val3 int
	fmt.Print("Enter three values: ")
	fmt.Scanf("%d%d%d", &val1, &val2, &val3)

	if val1 > val2 && val1 > val3 {
		fmt.Printf("%v is larger than %v and %v\n", val1, val2, val3)
	} else if val2 > val1 && val2 > val3 {
		fmt.Printf("%v is larger than %v and %v\n", val2, val1, val3)
	} else if val3 > val1 && val3 > val2 {
		fmt.Printf("%v is larger than %v and %v\n", val3, val1, val2)
	} else {
		fmt.Println("All values are equal")
	}
}
