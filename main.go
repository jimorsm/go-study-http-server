package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
	for k, v := range r.Header {
		r := ""
		for i, s := range v {
			r += s
			if i < len(v)-1 {
				r += ","
			}
			w.Header().Add(k, r)
		}
		fmt.Fprintf(w, "%q: %q\n", k, r)
		// w.Header().Set(k, r)
	}
}

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/healthz", HealthzHandler)
	http.ListenAndServe(":8000", nil)
}
