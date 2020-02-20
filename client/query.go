package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func executeQuery(url string) string {
	start := time.Now()
	response, _ := http.Get(url)
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("Request time: %d ms from url%s\n", time.Since(start).Nanoseconds()/time.Millisecond.Nanoseconds(), url)
	return fmt.Sprintf("%s from %s", body, url)
}
