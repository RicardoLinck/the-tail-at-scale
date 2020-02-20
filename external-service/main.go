package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/ishealthy", func(w http.ResponseWriter, r *http.Request) {
		rd := rand.New(rand.NewSource(time.Now().UnixNano()))
		requestPercentile := rd.Intn(100)
		waitTime := 0

		if requestPercentile > 96 {
			waitTime = 100
		}

		time.Sleep(time.Duration(waitTime+15) * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Healthy"))
	}).Methods(http.MethodGet)
	http.ListenAndServe(":8080", router)
}
