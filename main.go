package main

import (
	"errors"
	"fmt"
	"net/http"
)

type res struct {
	url    string
	result string
}

func main() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	c := make(chan res)
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

var errRequestsFailed = errors.New("Requests Failed")

func hitURL(url string, c chan<- res) {
	fmt.Println("Checking: ", url)
	result := "OK"
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		result = "FAILED"
	}
	c <- res{url: url, result: result}
}
