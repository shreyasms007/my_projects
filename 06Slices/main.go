package main

import "fmt"

func main() {

	fmt.Println("welcome to slices")

	fruitList := []string{"Apple", "Orange", "Banana"}

	fmt.Printf("The type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango")

	fmt.Println(fruitList)

}
