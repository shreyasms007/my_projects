package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("switch and case in golang")

	rand.Seed(time.Now().UnixMicro())

	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 , You can open")
	case 2:
		fmt.Println("Dice value is:2, You can move 2 steps")
	case 3:
		fmt.Printf("Dice value is:3 , You can move 3 steps .\n")
		fallthrough
	case 4:
		fmt.Printf("Dice value is:4 , You can move 4 steps .\n")
		fallthrough
	case 5:
		fmt.Printf("Dice value is:5 , You can move 5 steps .\n")
	case 6:
		fmt.Println("Dice value is:6 , You can roll the dice again.")
		fallthrough
	default:
		fmt.Printf("Dice vaue is : %v", diceNumber)
	}

}
