package main

import (
	"log"
	"log/slog"
	"os"
	"go_project/concurrency"
)

// 10:13 timestamp  https://www.youtube.com/watch?v=LvgVSSpwND8

// https://betterstack.com/community/guides/logging/logging-in-go/
func main() {
	log.Println("Hello from Go application!")

	defaultLogger := log.Default()
	defaultLogger.SetOutput(os.Stdout)
	log.Println("Hello from Go application!")

	logger := log.New(
		os.Stderr,
		"MyApplication: ",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC|log.Lshortfile,
	)

	logger.Println("Hello from Go application!")

	slog.Debug("Debug message")
	slog.Info("Info message")
	slog.Warn("Warning message")
	slog.Error("Error message")

	concurrency.RunConcurrency()
}
