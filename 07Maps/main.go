package main

import "fmt"

func main() {

	fmt.Println("Welcome to Maps in go ")

	laguages := make(map[int]string)

	laguages[3] = "Ruby"
	laguages[2] = "Python"
	laguages[1] = "JavaScript"

	fmt.Println("Languages in maps are : ", laguages)

	for key, value := range laguages {

		fmt.Printf("For value of : %v , Key is : %v.\n", value, key)
	}
}
