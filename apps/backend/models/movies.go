package models

import (
	"time"
)

type Movies struct {
	MovieId    int        `json:"movieId" gorm:"primaryKey"`
	MovieName  string     `json:"movieName"`
	UploadedBy string     `json:"uploadedBy"`
	UploadedOn time.Time  `json:"uploadedOn"`
}

type Comment struct {
	Comment string `json:"comment"`
}
