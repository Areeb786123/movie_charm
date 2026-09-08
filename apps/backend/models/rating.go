package models

type Rating struct {
	RatingID uint `json:"ratingId" gorm:"primaryKey"`
	MovieId  int  `json:"movieId"`
	Rating   int  `json:"rating"`
}

type MovieRatedEvent struct {
	EventID string `json:"eventId"`
	MovieID int    `json:"movieId"`
	Rating  int    `json:"rating"`
}
