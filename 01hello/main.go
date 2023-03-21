package main

import (
	"fmt"
)

func main() {

	name := []int{1, 2, 3, 4, 5, 6}
	for _, i := range name {

		if i%2 == 0 {
			fmt.Println(i, "is even number ")
		} else {
			fmt.Println(i, "is odd number")
		}

	}
}
