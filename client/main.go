package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	urls := []string{"http://localhost:8081/ishealthy", "http://localhost:8082/ishealthy", "http://localhost:8083/ishealthy"}

	router := mux.NewRouter()

	router.HandleFunc("/simple", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(urls[0])
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}).Methods(http.MethodGet)

	router.HandleFunc("/fanout", func(w http.ResponseWriter, r *http.Request) {
		result := queryFanOut(urls)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}).Methods(http.MethodGet)

	router.HandleFunc("/hedged", func(w http.ResponseWriter, r *http.Request) {
		result := queryWithHedgedRequestsWithContext(urls)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}).Methods(http.MethodGet)

	http.ListenAndServe(":8080", router)
}
