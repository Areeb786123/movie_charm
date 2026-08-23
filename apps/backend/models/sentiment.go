package models

type Sentiment struct {
	Message string `json:"message"`
	SentimentType string `json:"sentiment_type"`
}