package main

import "fmt"

func rev(s string) {

	var temp string

	for _, val := range s {

		temp = string(val) + temp
	}
	fmt.Println(temp)
}

func main() {

	n := "shreyas"

	rev(n)
}
