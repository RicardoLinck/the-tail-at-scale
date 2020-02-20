package main

import (
	"context"
	"time"
)

func queryWithHedgedRequestsWithContext(urls []string) string {
	ch := make(chan string, len(urls))
	ctx, cancel := context.WithCancel(context.Background())
	for _, url := range urls {
		go func(u string, c chan string) {
			c <- executeQueryWithContext(u, ctx)
		}(url, ch)

		select {
		case r := <-ch:
			cancel()
			return r
		case <-time.After(21 * time.Millisecond):
		}
	}

	return <-ch
}
