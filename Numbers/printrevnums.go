package main

import "fmt"

func main() {

	var n int
	fmt.Println("Enter the range: ")
	fmt.Scanln(&n)

	for i := n; i >= 0; i-- {
		fmt.Println(i)
	}
}
