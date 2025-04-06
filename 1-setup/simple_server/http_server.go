package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("server is running...")

	s := http.NewServeMux()
	s.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%+v\n", r)

		hName, err := os.Hostname()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("request received on host %s !\n", hName)))
	})

	sv := http.Server{
		Addr:    ":9999",
		Handler: s,
	}

	err := sv.ListenAndServe()
	panic(err)
}
