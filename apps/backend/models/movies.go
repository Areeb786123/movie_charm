package models

import (
	"time"
)

type Movies struct {
	MovieId       int       `json:"movieId" gorm:"primaryKey"`
	MovieName     string    `json:"movieName"`
	UploadedBy    string    `json:"uploadedBy"`
	UploadedOn    time.Time `json:"uploadedOn"`
	MovieLink     string    `json:"movieLink"`
	ImageUrl      string    `json:"imageUrl"`
	TrailorUrl    string    `json:"trailorUrl"`
	AverageRating float64   `json:"rating"`
	RatingCount   int       `json:"count"`
}

type Comment struct {
	Comment   string `json:"comment"`
	Sentiment string `json:"sentiment"`
	Message   string `json:"message"`
}
