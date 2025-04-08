package main

import (
	"github.com/fsnotify/fsnotify"
	"log/slog"
	"os"
)

func main() {
	logFilePath := os.Args[1]
	if len(logFilePath) == 0 {
		panic("log file path is empty")
	}

	slog.Info("log file path", "path", logFilePath)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		slog.Error(err.Error(), "method", "create watcher")
		os.Exit(1)
	}
	defer watcher.Close()

	if err := watcher.Add(logFilePath); err != nil {
		slog.Error(err.Error(), "method", "add watcher")
		os.Exit(1)
	}

	slog.Info("log collector started")

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				slog.Info("file modified", "path", event.Name, "op", event.Op)
			}
			continue
		case err := <-watcher.Errors:
			slog.Error(err.Error(), "method", "add watcher")
		}
	}
}
