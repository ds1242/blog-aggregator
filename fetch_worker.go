package main

import (
	// "sync"
	"time"
	"fmt"
)


func FetchFeedFromDB(numberOfFeeds int) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(60 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
}

