package httpapi

import (
	"fmt"
	"html"
	"net/http"
)

func Build(mux *http.ServeMux) {
	mux.HandleFunc("/", hello)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
