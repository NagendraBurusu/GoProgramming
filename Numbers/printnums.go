package main

import "fmt"

func main() {

	var n int
	fmt.Println("Enter the range :")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}

}
