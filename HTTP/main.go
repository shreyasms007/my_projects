package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://instagram.com",
	}

	for _, link := range links {
		time.Sleep(1 * time.Second)
		go getLinks(link)
	}

}

func getLinks(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		return
	}
	fmt.Println(link, "is up")
}
