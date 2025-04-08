package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logFilePath := os.Args[1]
	if len(logFilePath) == 0 {
		panic("log file path is empty")
	}

	address := os.Args[2]
	if len(address) == 0 {
		panic("address is empty")
	}

	log.Println("log file path:", logFilePath)
	log.Println("address:", address)

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	fileHandler := slog.NewJSONHandler(logFile, nil)
	logger := slog.New(fileHandler)
	slog.SetDefault(logger)

	router := http.NewServeMux()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		logger.Info(request.Method+" "+request.RequestURI, "key", "value")

		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte("got your request\n"))
	})

	server := &http.Server{
		Addr:    address,
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		logger.Error(err.Error(), "method", "listen&server")
	}
}
