package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
	for k, v := range r.Header {
		r := ""
		for i, s := range v {
			r += s
			if i < len(v)-1 {
				r += ","
			}

		}
		fmt.Fprintf(w, "%q: %q\n", k, r)
		w.Header().Set(k, r)
	}
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8000", nil)
}
