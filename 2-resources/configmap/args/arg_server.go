package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println(os.Args)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("got request for:", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Hello World"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
