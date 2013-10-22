// This program demostrates how to read an rss feed and
// display the result

package main

import (
	"github.com/goinggo/gomiami/meetup-02/6-rss/helper"
	"github.com/goinggo/gomiami/meetup-02/6-rss/rss"
)

const (
	FEED = "http://rss.cnn.com/rss/cnn_topstories.rss"
)

// main is the entry point for the program
func main() {
	helper.WriteStdout("main", "main", "main", "Started")

	// Retrieve the RSS feed document
	rssDocument, err := rss.RetrieveRssFeed("main", FEED)

	if err != nil {
		helper.WriteStdoutf("main", "main", "main", "Completed : ERROR : %s", err)
		return
	}

	// Display the results
	rss.DisplayRssFeed("main", rssDocument)

	helper.WriteStdout("main", "main", "main", "Completed")
}
