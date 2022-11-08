package main

import (
	"fmt"
	"net/http"
	"time"
)

type linkStatus struct {
	link         string
	status       string
	currentCount int64
}

var countOfRequestsMade int64

func main() {
	countOfRequestsMade = 0
	listOfUrls := []string{
		"https://praful.netlify.app",
		"https://www.google.com",
		"https://www.facebook.com",
	}

	c := make(chan linkStatus)

	for _, link := range listOfUrls {

		go checkLink(link, c)

	}

	// for {
	// 	ls := <-c
	// 	fmt.Println(ls.status)
	// 	fmt.Println("Number of requests made : ", ls.currentCount)
	// 	go checkLink(ls.link, c)
	// }

	for recivedLink := range c {
		ls := recivedLink
		// fmt.Println(ls.status)
		// fmt.Println("Number of requests made : ", ls.currentCount)
		go func(lsc linkStatus) {
			time.Sleep(time.Second * 3)
			go checkLink(lsc.link, c)

		}(ls)
	}

}

func checkLink(link string, c chan linkStatus) {
	l := linkStatus{}

	countOfRequestsMade++
	l.currentCount = countOfRequestsMade
	_, err := http.Get(link)

	l.link = link

	if err != nil {
		fmt.Println(link, "Is down")
		l.status = "Yep its Down"
		c <- l
		return
	}

	fmt.Println(link, " Is working fine")
	l.status = "Yep its Up!"
	c <- l

}
