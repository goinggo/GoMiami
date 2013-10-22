// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rss

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/goinggo/gomiami/meetup-02/6-rss/helper"
	"io/ioutil"
	"net/http"
)

//** NEW TYPES

// _RSSItem defines the fields associated with the item tag in the buoy RSS document
type RSSItem struct {
	XMLName     xml.Name `xml:"item"`
	PubDate     string   `xml:"pubDate"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	GUID        string   `xml:"guid"`
	GeoRssPoint string   `xml:"georss:point"`
}

// _RSSImage defines the fields associated with the image tag in the buoy RSS document
type RSSImage struct {
	XMLName xml.Name `xml:"image"`
	Url     string   `xml:"url"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
}

// _RSSChannel defines the fields associated with the channel tag in the buoy RSS document
type RSSChannel struct {
	XMLName        xml.Name  `xml:"channel"`
	Title          string    `xml:"title"`
	Description    string    `xml:"description"`
	Link           string    `xml:"link"`
	PubDate        string    `xml:"pubDate"`
	LastBuildDate  string    `xml:"lastBuildDate"`
	TTL            string    `xml:"ttl"`
	Language       string    `xml:"language"`
	ManagingEditor string    `xml:"managingEditor"`
	WebMaster      string    `xml:"webMaster"`
	Image          RSSImage  `xml:"image"`
	Items          []RSSItem `xml:"item"`
}

// _RSSDocument defines the fields associated with the buoy RSS document
type RSSDocument struct {
	XMLName xml.Name   `xml:"rss"`
	Channel RSSChannel `xml:"channel"`
	Uri     string
}

//** PUBLIC FUNCTIONS

// RetrieveRssFeed performs a HTTP Get request to the RSS feed and serializes the results
//  goRoutine: The Go routine making the call
//  uri: The uri of the rss feed to retrieve
func RetrieveRssFeed(goRoutine string, uri string) (rssDocument *RSSDocument, err error) {
	defer helper.CatchPanic(&err, goRoutine, "rss", "RetrieveRssFeed")

	helper.WriteStdoutf(goRoutine, "rss", "RetrieveRssFeed", "Started : Uri[%s]", uri)

	if uri == "" {
		err = errors.New("No RSS Feed Uri Provided")
		helper.WriteStdoutf(goRoutine, "rss", "RetrieveRssFeed", "Completed : ERROR : %s", err)
		return rssDocument, err
	}

	resp, err := http.Get(uri)

	if err != nil {
		helper.WriteStdoutf(goRoutine, "rss", "RetrieveRssFeed", "ERROR - Completed : HTTP Get : %s : %s", uri, err)
		return rssDocument, err
	}

	defer func() {
		resp.Body.Close()
		helper.WriteStdoutf(goRoutine, "rss", "RetrieveRssFeed", "Defer Completed : Uri[%s]", uri)
	}()

	rawDocument, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		helper.WriteStdoutf(goRoutine, "rss", "RetrieveRssFeed", "ERROR - Completed : Read Resp : %s : %s", uri, err)
		return rssDocument, err
	}

	decoder := xml.NewDecoder(bytes.NewReader(rawDocument))

	rssDocument = &RSSDocument{}
	err = decoder.Decode(rssDocument)

	if err != nil {
		helper.WriteStdoutf(goRoutine, "rss", "RetrieveRssFeed", "ERROR - Completed : Decode : %s : %s", uri, err)
		return rssDocument, err
	}

	// Save the uri to the feed
	rssDocument.Uri = uri

	helper.WriteStdoutf(goRoutine, "rss", "RetrieveRssFeed", "Completed : Uri[%s] Title[%s]", uri, rssDocument.Channel.Title)

	return rssDocument, err
}

// DisplayRssFeed display the specified feed
//  goRoutine: The Go routine making the call
//  rssDocument: The document to display
func DisplayRssFeed(goRoutine string, rssDocument *RSSDocument) {
	defer helper.CatchPanic(nil, goRoutine, "rss", "DisplayRssFeed")

	fmt.Printf("\n\nURL: %s\n\n", rssDocument.Uri)

	fmt.Printf("Title: %s\n", rssDocument.Channel.Title)
	fmt.Printf("PubDate: %s\n", rssDocument.Channel.PubDate)
	fmt.Printf("Description: %s\n\n", rssDocument.Channel.Description)

	for _, item := range rssDocument.Channel.Items {

		fmt.Printf("Title: %s\n", item.Title)
		fmt.Printf("PubDate: %s\n", item.PubDate)
		fmt.Printf("Description: %s\n\n", item.Description)
	}
}
