package main

import "fmt"

func main() {

	var n int
	fmt.Println("Enter the Range :")
	fmt.Scanln(&n)
	fmt.Printf("Prime Number between 1 and %d", n)
	for i := 1; i < n; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(i)
		}
	}
}
