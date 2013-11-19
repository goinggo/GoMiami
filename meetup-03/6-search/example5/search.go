/*
	Example: Google Search 3.0

	Given a query, return a page of search results (and some ads).
	Send the query to web search, image search, YouTube, Maps, News, etc. then mix the results.

	No locks. No condition variables. No callbacks

	Reduce tail latency using replicated search servers

	Run the same search against multiple servers in their own Goroutine but only return searches
	that complete in 80 Milliseconds or less

	All three searches SHOULD always come back in under 80 milliseconds
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web1   = fakeSearch("web")
	Web2   = fakeSearch("web")
	Image1 = fakeSearch("image")
	Image2 = fakeSearch("image")
	Video1 = fakeSearch("video")
	Video2 = fakeSearch("video")
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
		c <- First(query, Web1, Web2)
	}()

	go func() {
		c <- First(query, Image1, Image2)
	}()

	go func() {
		c <- First(query, Video1, Video2)
	}()

	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}

	return results
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)

	// Define a function that takes the index to the replica function to use.
	// Then it executes that function writing the results to the channel
	searchReplica := func(i int) {
		c <- replicas[i](query)
	}

	// Run each replica function in its own Goroutine
	for i := range replicas {
		go searchReplica(i)
	}

	// As soon as one of the replica functions write a result, return
	return <-c
}
