// Program demostrates how to read an rss feed and
// display the result.
package main

import (
	"github.com/goinggo/gomiami/meetup-02/6-rss/helper"
	"github.com/goinggo/gomiami/meetup-02/6-rss/rss"
)

const (
	// Feed contains the URL to the feed to retrieve and process.
	Feed = "http://rss.cnn.com/rss/cnn_topstories.rss"
)

// main is the entry point for the program.
func main() {
	helper.WriteStdout("main", "main", "main", "Started")

	// Retrieve the RSS feed document.
	document, err := rss.RetrieveFeed("main", Feed)
	if err != nil {
		helper.WriteStdoutf("main", "main", "main", "Completed : ERROR : %s", err)
		return
	}

	// Display the results.
	rss.DisplayFeed("main", document)

	helper.WriteStdout("main", "main", "main", "Completed")
}
