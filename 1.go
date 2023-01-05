package main

import "fmt"

func main() {

	var n, a, r int

	fmt.Scanln(&n, &a, &r)

	fmt.Print(0)
	for i := 1; i <= n; i++ {
		fmt.Printf(" + %d", a*i*r)
	}
	fmt.Println()

}