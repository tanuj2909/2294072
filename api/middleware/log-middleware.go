package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	logAPI = "http://20.244.56.144/evaluation-service/logs"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		log.Printf("Incoming Request: %s %s", c.Method(), c.Path())

		err := c.Next()

		duration := time.Since(start)

		if err != nil {
			Log("backend", "error", "route", fmt.Sprint("%v", err))
		} else {
			message := fmt.Sprintf("Request %s %s took %v and responded with status %d",
				c.Method(),
				c.Path(),
				duration,
				c.Response().StatusCode(),
			)
			Log("backend", "info", "route", message)
		}

		return err
	}
}

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
