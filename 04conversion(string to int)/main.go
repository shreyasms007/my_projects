package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to our pizza")
	fmt.Println("Please rate our pizza from 1 to 5 : ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for Rating : ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("added 1 to your number : ", numRating+1)
	}

}
