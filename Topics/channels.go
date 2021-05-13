package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Channels Example")
	links := []string{
		"http://www.google.com",
		"http://www.spotify.com",
		"http://www.amazon.com",
	}
	c := make(chan string)
	for _, link := range links {
		go printExample(c, link)
		//No sleep statements in main routine since it will take time to recieve from channel
	}
	// ....Recieve number of messages equal to the number of requests we make
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }
	// ....ReFetching a link after it has been successfully fetched either with or wihtout error
	// 	for {
	// 		go printExample(c, <-c)
	// 	}
	// Best way to use sleep with function literal(anonymous)
	for l := range c {
		go func(link string) {
			time.Sleep((time.Second * 2))
			printExample(c, link)
		}(l)
	}
}

func printExample(c chan string, link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("exists error")
		c <- link
		return
	}
	fmt.Println(link, "is out")
	c <- link
}
