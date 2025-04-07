package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync/atomic"
)

func main() {
	slog.Info("starting http server", "port", "8989")

	var count atomic.Int32

	h := http.NewServeMux()

	h.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("got healthcheck request", "request", fmt.Sprintf("%+v", r))

		// response unhealthy after 3 healthcheck probe
		if count.Load() > 3 {
			slog.Error("healthcheck failed...")
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("healthcheck failed\n"))
			return
		}

		slog.Info("healthcheck probe...")
		count.Add(1)
		return
	})

	httpServer := http.Server{
		Addr:    "0.0.0.0:8989",
		Handler: h,
	}

	slog.Info("starting http server", "address", httpServer.Addr)

	err := httpServer.ListenAndServe()
	if err != nil {
		slog.Error(err.Error(), "cause", "server-failure-on-listen")
	}
}
