package main

import (
	"encoding/xml"
	"net/http"
	"io"
)


func FetchRSSFeed(url string) (RSS, error ){
	resp, err := http.Get(url)
	if err != nil {
		return RSS{}, err
	}
	defer resp.Body.Close()
	// // Read the response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return RSS{}, err
    }
	// Parse the XML
    var rss RSS
    err = xml.Unmarshal(body, &rss)
    if err != nil {
        return RSS{}, err
    }

    return rss, nil

}