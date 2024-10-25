package rest

import (
	"bytes"
	"fmt"
	"net/http"
)

func SendMessage(message string) error {
	req, err := http.NewRequest("POST", "http://localhost:3001/api/v1/message", bytes.NewBuffer([]byte(message)))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 response: %d", resp.StatusCode)
	}

	return nil
}
