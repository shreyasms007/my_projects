package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://LCO.dev"

func main() {

	response, err := http.Get(url)

	if err != nil {

		panic(err)
	}

	fmt.Println(response)

	fmt.Printf("Response is of type : %T\n", response)

	databytes, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(databytes))

}
