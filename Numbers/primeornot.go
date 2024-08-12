package main

import "fmt"

func main() {
	var n, i int
	fmt.Println("Enter the number :")
	fmt.Scanln(&n)
	if n == 1 {
		fmt.Println("Prime starts from 2")
	}
	for i = 2; i < n; i++ {
		if n%i == 0 {
			fmt.Printf("%d not a prime", n)
			break
		}
	}
	if n == i {
		fmt.Printf("%d is a prime number", n)
	}
}
