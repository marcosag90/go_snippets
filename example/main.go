package main

import (
	"fmt"
	"net/http"
)


func main() {
	links := []string {
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://www.amazon.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go checkLink(l, c)
	}
	fmt.Println(<-c)
}

func checkLink( link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "may be down")
		c <- link
		return
	}
	c <- link
	fmt.Println(link, "is up")
}
