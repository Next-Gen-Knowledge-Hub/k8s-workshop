package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		slog.Error("error on reading port from environment")
		os.Exit(1)
	}

	stage := os.Getenv("STAGE")
	if len(stage) == 0 {
		slog.Error("error on reading stage from environment")
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("got http request", "method", r.Method, "host", r.Host)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprintf(`{"message": "Hello From Env Server", "stage": "%s"}`, stage)))
	})

	slog.Info("server is running", "port", port, "stage", stage)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		slog.Error("error on starting http server", "error", err)
		os.Exit(1)
	}
}
