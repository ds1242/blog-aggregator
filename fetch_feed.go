package main

import (
	"encoding/xml"
	"net/http"
	"io"
)

type RSS struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title 	string	`xml:"title"`
	Items	[]Item	`xml:"item"`
}

type Item struct {
	Title	string	`xml:"title"`
	Link	string	`xml:"link"`
}


func FetchFeed(url string) (RSS, error ){
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