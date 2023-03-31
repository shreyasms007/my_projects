package main

import "fmt"

func main() {

	sum := proAdder(2, 3, 4, 5)

	greeter()

	fmt.Println(sum)
}

func proAdder(values ...int) int{

	total := 0
	for _, val := range values {
		total += val
	}
	return total,

}

func greeter() {
	fmt.Printf("welcome to golang and the value is : %v", proAdder())
}
