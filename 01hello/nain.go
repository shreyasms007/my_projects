package main

import "fmt"

func main() {

	a, b := reverse(3, 4)
	fmt.Println(a, b)

}

func reverse(x, y int) (int, int) {
	return y, x

}
