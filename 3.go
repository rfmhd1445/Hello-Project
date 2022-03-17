package main

import "fmt"

func main() {

	var n, bil int

	fmt.Scanln(&n)

	var i int = 0
	var total = 0
	for i < n {
		fmt.Scanln(&bil)
		for bil < 0 || bil > 9 {
			fmt.Scanln(&bil)
		}
		total += bil
		i++
	}

	fmt.Println(total)

}
