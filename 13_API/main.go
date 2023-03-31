package main

import "fmt"

var veg []string

func main() {

	example()

}

func example() {
	veg = []string{"apples", "mango", "banana", "watermelon"}
	for _, val := range veg {
		if val == "apples" {
			fmt.Println("woo apples")
		}

		fmt.Println(val)
	}

}
