package models

type Comments struct {
	CommentID uint   `json:"commentId" gorm:"primaryKey"`
	MovieId   int    `json:"movieId"`
	Comment   string `json:"comment"`
}
