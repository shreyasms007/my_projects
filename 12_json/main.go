package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	encodeJson()

}

func encodeJson() {

	lcocourses := []course{
		{"REACTJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "json"}},
		{"PYTHON Bootcamp", 199, "LearnCodeOnline.in", "BCD123", []string{"full stack", "json"}},
		{"MERN", 299, "LearnCodeOnline.in", "lkj123", nil},
	}

	finalJason, err := json.MarshalIndent(lcocourses, "lco", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJason)

}

// func decodeJason() {

// 	jsonDataFromWeb := []byte(` {
// 		            "courseName": "REACTJS Bootcamp",
// 		            "Price": 299,
// 		            "website": "LearnCodeOnline.in",
// 		            "tags": [ "web-dev", "json"]
// 		    },
// 			`)

// 			var lcoCourse

// }
