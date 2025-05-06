package main

import (
	"fmt"
	"go-router-rtoolkit/router"
	"net/http"
)

func main() {

	r := router.NewRouter()
	r.HandleFunc("GET", "/test/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := router.Param(r, "name")
		fmt.Fprintf(w, "Hello test %s", name)
	})
	r.HandleFunc("GET", "/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := router.Param(r, "name")
		fmt.Fprintf(w, "change %s", name)
	})

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", r)

}
