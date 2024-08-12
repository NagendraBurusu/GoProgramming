package main

import "fmt"

func main() {

	var n int
	fmt.Println("Enter the numbers :")
	fmt.Scanln(&n)

	if n%2 == 0 {
		fmt.Printf("%d is a even number.", n)
	} else {
		fmt.Printf("%d is a odd number.", n)
	}
}
