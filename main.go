package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello , go router rtookit")

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello test in go router")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
