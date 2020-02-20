package main

func queryFanOut(urls []string) string {
	ch := make(chan string, len(urls))
	for _, url := range urls {
		go func(u string) {
			ch <- executeQuery(u)
		}(url)
	}
	return <-ch
}
