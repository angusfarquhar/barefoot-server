package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	salstr := r.URL.Path[1:]
	fmt.Fprintf(w, "Your salary is %s.\n", salstr)

	salInt, _ := strconv.Atoi(salstr)

	buckets := makeBarefootBuckets(salInt)
	fmt.Fprintf(w, "SPLURGE bucket is %v,\n", buckets["splurge"])
	fmt.Fprintf(w, "SMILE bucket is %v,\n", buckets["smile"])
	fmt.Fprintf(w, "FIRE EXTINGUISHER bucket is %v.\n", buckets["fire extinguisher"])

}

func makeBarefootBuckets(salary int) map[string]int {
	buckets := make(map[string]int)
	buckets["splurge"] += salary / 10
	buckets["smile"] += salary / 10
	buckets["fire extinguisher"] += salary / 5

	return buckets
}

func main() {
	//
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
