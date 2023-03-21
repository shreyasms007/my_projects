package main

import "fmt"

type fruit []string

func main() {

	market := fruit{}

	marketFruits := []string{"Apple", "Orange", "Pineapple", "grapes"}
	marketPrice := []string{"10.rs", "25.rs", "30.rs", "45.rs"}

	for _, Fruits := range marketFruits {
		for _, Price := range marketPrice {
			market := append(market, Fruits+" price is "+Price)
			fmt.Println(market)
		}
	}

}
