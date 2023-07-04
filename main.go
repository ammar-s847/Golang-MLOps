package main

import (
	"fmt"
	"log"
	"net/http"
)

func mlOps(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mlops Enpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", mlOps)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func main() {
	handleRequests()
}
