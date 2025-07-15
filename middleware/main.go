package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const (
	logAPI = "http://20.244.56.144/evaluation-service/logs"
)

type LogRequest struct {
	Stack   string `json:"stack"`
	Level   string `json:"level"`
	Package string `json:"package"`
	Message string `json:"message"`
}

func Log(stack, level, pkg, message string) {
	accessToken := os.Getenv("ACCESS_TOKEN")
	logEntry := LogRequest{
		Stack:   stack,
		Level:   level,
		Package: pkg,
		Message: message,
	}

	payload, err := json.Marshal(logEntry)
	if err != nil {
		fmt.Printf("Failed to marshal log entry: %v\n", err)
		return
	}

	client := &http.Client{Timeout: 5 * time.Second}

	req, err := http.NewRequest("POST", logAPI, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Printf("Failed to create log request: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to send log: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Log API returned status: %s\n", resp.Status)
	} else {
		fmt.Println("Log sent successfully.")
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	Log("backend", "error", "handler", "received string, expected bool")
	Log("backend", "fatal", "db", "Critical database connection failure.")
}
