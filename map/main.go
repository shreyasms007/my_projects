package main

import "fmt"

func main() {

	n := 254

	if n%3 == 0 && n%5 == 0 {
		fmt.Println("The number is divisible by 3 and 5")
	} else if n%3 == 0 && n%5 != 0 {

		fmt.Println("The number is divisible by 3 not by 5")

	} else if n%3 != 0 && n%5 == 0 {
		fmt.Println("The number is divisible by 5 and not 3")
	} else {
		fmt.Println("The number is not divisible by both 3 and 5 ")
	}

}
