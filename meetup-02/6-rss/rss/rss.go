// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rss implements support for retrieving and reading rss feeds.
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

// Item defines the fields associated with the item tag in the buoy RSS document.
type Item struct {
	XMLName     xml.Name `xml:"item"`
	PubDate     string   `xml:"pubDate"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	GUID        string   `xml:"guid"`
	GeoRssPoint string   `xml:"georss:point"`
}

// Image defines the fields associated with the image tag in the buoy RSS document.
type Image struct {
	XMLName xml.Name `xml:"image"`
	URL     string   `xml:"url"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
}

// Channel defines the fields associated with the channel tag in the buoy RSS document.
type Channel struct {
	XMLName        xml.Name `xml:"channel"`
	Title          string   `xml:"title"`
	Description    string   `xml:"description"`
	Link           string   `xml:"link"`
	PubDate        string   `xml:"pubDate"`
	LastBuildDate  string   `xml:"lastBuildDate"`
	TTL            string   `xml:"ttl"`
	Language       string   `xml:"language"`
	ManagingEditor string   `xml:"managingEditor"`
	WebMaster      string   `xml:"webMaster"`
	Image          Image    `xml:"image"`
	Items          []Item   `xml:"item"`
}

// Document defines the fields associated with the buoy RSS document.
type Document struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
	URI     string
}

//** PUBLIC FUNCTIONS

// RetrieveFeed performs a HTTP Get request to the RSS feed and serializes the results
func RetrieveFeed(goRoutine string, uri string) (*Document, error) {
	helper.WriteStdoutf(goRoutine, "rss", "RetrieveFeed", "Started : URI[%s]", uri)

	if uri == "" {
		err := errors.New("No RSS Feed URI Provided")
		helper.WriteStdoutf(goRoutine, "rss", "RetrieveFeed", "Completed : ERROR : %s", err)
		return nil, err
	}

	resp, err := http.Get(uri)
	if err != nil {
		helper.WriteStdoutf(goRoutine, "rss", "RetrieveFeed", "ERROR - Completed : HTTP Get : %s : %s", uri, err)
		return nil, err
	}

	defer resp.Body.Close()

	rawDocument, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		helper.WriteStdoutf(goRoutine, "rss", "RetrieveFeed", "ERROR - Completed : Read Resp : %s : %s", uri, err)
		return nil, err
	}

	var document Document
	err = xml.NewDecoder(bytes.NewReader(rawDocument)).Decode(&document)
	if err != nil {
		helper.WriteStdoutf(goRoutine, "rss", "RetrieveFeed", "ERROR - Completed : Decode : %s : %s", uri, err)
		return nil, err
	}

	// Save the uri to the feed
	document.URI = uri

	helper.WriteStdoutf(goRoutine, "rss", "RetrieveFeed", "Completed : URI[%s] Title[%s]", uri, document.Channel.Title)
	return &document, nil
}

// DisplayFeed display the specified feed
func DisplayFeed(goRoutine string, document *Document) {
	fmt.Printf("\n\nURL: %s\n\n", document.URI)

	fmt.Printf("Title: %s\n", document.Channel.Title)
	fmt.Printf("PubDate: %s\n", document.Channel.PubDate)
	fmt.Printf("Description: %s\n\n", document.Channel.Description)

	for _, item := range document.Channel.Items {
		fmt.Printf("Title: %s\n", item.Title)
		fmt.Printf("PubDate: %s\n", item.PubDate)
		fmt.Printf("Description: %s\n\n", item.Description)
	}
}
