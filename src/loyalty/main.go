// Loyalty points: checkout earns points for an order; the header shows a shopper's balance.
package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/points", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("0")) })
	http.HandleFunc("/earn", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusAccepted) })
	log.Fatal(http.ListenAndServe(":"+os.Getenv("LOYALTY_PORT"), nil))
}
