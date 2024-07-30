package main

import (
	"encoding/xml"
	"net/http"
	"io"
)


func fetchRSSFeed(url string) (*RSSFeed, error ){
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// // Read the response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
	// Parse the XML
    var rss RSSFeed
    err = xml.Unmarshal(body, &rss)
    if err != nil {
        return nil, err
    }

    return &rss, nil

}