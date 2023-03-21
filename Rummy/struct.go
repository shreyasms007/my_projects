package main

import "fmt"

type contactInfo struct {
	gmail   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "stewart",
		contactInfo: contactInfo{
			gmail:   "alex@gmail.com",
			zipcode: 573103,
		},
	}
	alex.print()
	alex.updateName("alexy")

}

func (p person) print() {
	fmt.Println(p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName

}
