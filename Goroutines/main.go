package main

import (
	"fmt"
)

func main() {
	// go greeter("hello")

	// greeter("world")

}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(1 * time.Second)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {

	res, err := http.get(endpoint)

	if err != nil {
		fmt.Println("OOPS an error ")

	}
	fmt.Println("%dstatus code for %s", res.StatusCode, endpoint)

}
