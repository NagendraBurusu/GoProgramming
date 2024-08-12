package main

import "fmt"

func main() {

	var n, sum int
	fmt.Println("Enter the range: ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Printf("Sum of %d range is %d", n, sum)

}
