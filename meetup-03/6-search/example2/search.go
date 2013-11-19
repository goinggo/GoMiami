/*
	Example: Google Search

	Given a query, return a page of search results (and some ads).
	Send the query to web search, image search, YouTube, Maps, News, etc. then mix the results.

	Run the Web, Image and Video searches concurrently, and wait for all results.
	No locks. No condition variables. No callbacks

	Run each search in their own Goroutine and wait for all searches to complete before display results
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web   = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type (
	Result string
	Search func(query string) Result
)

func main() {
	rand.Seed(time.Now().UnixNano())

	start := time.Now()
	results := Google("golang")
	elasped := time.Since(start)

	fmt.Println(results)
	fmt.Println(elasped)
}

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func Google(query string) (results []Result) {
	c := make(chan Result)

	go func() {
		c <- Web(query)
	}()

	go func() {
		c <- Image(query)
	}()

	go func() {
		c <- Video(query)
	}()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}

	return results
}
