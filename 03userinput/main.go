package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to golang"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ener the rating for our pizza: ")

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("error")
	}

	fmt.Println("Thanks for rating :", input)
	fmt.Printf("Type of this rating is %T\n", input)

}
