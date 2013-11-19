/*
	Example: Google Search 2.1 - Avoid Timeouts

	Given a query, return a page of search results (and some ads).
	Send the query to web search, image search, YouTube, Maps, News, etc. then mix the results.

	No locks. No condition variables. No callbacks

	Replicate the servers. Send requests to multiple replicas, and use the first response.

	Run the same search against multiple servers in their own Goroutine
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	Result string
	Search func(query string) Result
)

func main() {
	rand.Seed(time.Now().UnixNano())

	start := time.Now()

	// Run the search against two replicas
	result := First("golang",
		fakeSearch("replica 1"),
		fakeSearch("replica 2"))

	elasped := time.Since(start)

	fmt.Println(result)
	fmt.Println(elasped)
}

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
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
