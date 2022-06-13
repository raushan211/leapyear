package main

import "fmt"

func main() {
	fmt.Println("Enter Year")
	var year int
	fmt.Scanln(&year)

	if year%400 == 0 && year%4 == 0 {
		fmt.Println("Leap Year")
	} else {
		fmt.Println("Not Leap Year")

	}
}
