package kestra

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Kestra struct {
	Logger *KestraLogger
}

func NewKestra() *Kestra {
	return &Kestra{}
}

type KestraLogger struct{}

type LogEntry struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

type LogOutput struct {
	Logs []LogEntry `json:"logs"`
}

func (k *KestraLogger) logMessage(level, message string) {
	timestamp := time.Now().Format(time.RFC3339Nano)
	logEntry := LogEntry{
		Level:   level,
		Message: timestamp + " - " + message,
	}

	logOutput := LogOutput{
		Logs: []LogEntry{logEntry},
	}

	jsonOutput, err := json.Marshal(logOutput)
	if err != nil {
		log.Fatalf("Error marshalling log output: %v", err)
	}

	os.Stdout.Write([]byte("::" + string(jsonOutput) + "::\n"))
}

func (k *KestraLogger) Trace(message string) {
	k.logMessage("TRACE", message)
}

func (k *KestraLogger) Debug(message string) {
	k.logMessage("DEBUG", message)
}

func (k *KestraLogger) Info(message string) {
	k.logMessage("INFO", message)
}

func (k *KestraLogger) Warning(message string) {
	k.logMessage("WARNING", message)
}

func (k *KestraLogger) Error(message string) {
	k.logMessage("ERROR", message)
}

func (k *KestraLogger) Critical(message string) {
	k.logMessage("CRITICAL", message)
}
