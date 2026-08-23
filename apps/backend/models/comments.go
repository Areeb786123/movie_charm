package models

type Comments struct {
	CommentID uint   `json:"commentId" gorm:"primaryKey"`
	MovieId   int    `json:"movieId"`
	Comment   string `json:"comment"`
	Sentiment string `json:"sentiment"`
	Message   string `json:"message"`
}

type CommentResponse struct {
	Message   string `json:"message"`
	Sentiment string `json:"sentiment"`
}
