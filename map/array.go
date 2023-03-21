package main

import "fmt"

func main() {

	var revnum, remainder int

	fmt.Print("Enter the Number to Reverse = ")
	fmt.Scanln(&revnum)

	reverse := 0

	for revnum > 0 {
		remainder = revnum % 10
		reverse = reverse*10 + remainder
		revnum = revnum / 10
	}

	fmt.Println("The Reverse of the Given Number = ", reverse)
}
