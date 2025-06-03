package kestra

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type outputsPayload struct {
	Outputs map[string]interface{} `json:"outputs"`
}

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

func (k *KestraLogger) Warn(message string) {
	k.logMessage("WARN", message)
}

func (k *KestraLogger) Error(message string) {
	k.logMessage("ERROR", message)
}

func Outputs(outputs map[string]interface{}) {
	if outputs == nil {
		log.Println("Warning: Outputs called with nil map.")
		return
	}
	payload := outputsPayload{Outputs: outputs}
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling outputs to JSON: %v\n", err)
		return
	}
	fmt.Printf("::%s::\n", string(jsonBytes))
}
