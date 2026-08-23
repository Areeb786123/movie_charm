package service

import (
	"backend/dto"
	"backend/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendSentiment(sentiment string) (*models.Sentiment, error) {
	requestBody := dto.AIRequest{
		Comment: sentiment,
	}
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	// this will call another  micro service to get the sentiment from the ai model
	resp, err := http.Post(
		"http://ai-service:8000/analyze",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("AI service returned status: %d", resp.StatusCode)
	}

	var result models.Sentiment

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
