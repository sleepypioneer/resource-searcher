package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

func main() {
	fs := http.FileServer(http.Dir("articles"))
	http.Handle("/articles/", http.StripPrefix("/articles/", fs))

	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
