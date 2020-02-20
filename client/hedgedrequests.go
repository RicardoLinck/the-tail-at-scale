package main

import (
	"time"
)

func queryWithHedgedRequests(urls []string) string {
	ch := make(chan string, len(urls))
	for _, url := range urls {
		go func(u string, c chan string) {
			c <- executeQuery(u)
		}(url, ch)

		select {
		case r := <-ch:
			return r
		case <-time.After(21 * time.Millisecond):
		}
	}

	return <-ch
}
