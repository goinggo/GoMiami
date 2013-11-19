/*
	Example: Google Search

	Given a query, return a page of search results (and some ads).
	Send the query to web search, image search, YouTube, Maps, News, etc. then mix the results.

	Google function takes a query and returns a slice of Results (which are just strings)
	Google invokes Web, Image and Video searches serially, appending them to the results slice.

	Run each search in series
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
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))

	return results
}
