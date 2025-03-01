package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/Tanmoy095/Toll-Calculator.git/types"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "the listen address of the HTTP server")
	flag.Parse()

	var (
		store = NewMemoryStore()
		svc   = NewInvoiceAggregator(store)
	)
	svc = NewLoggingMiddleware(svc)

	fmt.Println("this is working fine")
	makeHttpTransport(*listenAddr, svc)

}
func makeHttpTransport(listenAddr string, svc Aggregator) {
	fmt.Println("HTTP transport running on port ", listenAddr)
	http.HandleFunc("/aggregate", handleAggregate(svc))
	http.ListenAndServe(listenAddr, nil)

}

func handleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//read request
		//call service
		//write response
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {

			writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})

			return
		}
		if err := svc.AggregateDistance(distance); err != nil {
			writeJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

	}

}

func writeJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
