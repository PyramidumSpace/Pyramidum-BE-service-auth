package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	fmt.Println("Starting service-auth")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.Header.Get("User-Agent")))
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		return
	}
}
