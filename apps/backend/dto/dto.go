package dto

type CreateMovieRequest struct {
	MovieName  string `json:"movieName" binding:"required"`
	UploadedBy string `json:"uploadedBy" binding:"required"`
}

type CreateComments struct {
	MovieID int    `json:"movieID"`
	Comment string `json:"comment" binding:"required"`
}
