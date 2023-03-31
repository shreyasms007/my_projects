package main

import "fmt"

func main() {

	num := []int{546, 876, 233, 921, 934, 94, 987, 5647}

	for _, value := range num {

		if value%2 == 0 {
			fmt.Println("even number is : ", value)
		} else {
			fmt.Println("odd numer is : ", value)
		}
	}
}
