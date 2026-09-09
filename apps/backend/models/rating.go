package models

type Rating struct {
	RatingID uint `json:"ratingId" gorm:"column:id;primaryKey"`
	MovieId  int  `json:"movieId"`
	Rating   int  `json:"rating"`
}

func (Rating) TableName() string {
	return "rating"
}

type MovieRatedEvent struct {
	EventID string `json:"eventId"`
	MovieID int    `json:"movieId"`
	Rating  int    `json:"rating"`
}
