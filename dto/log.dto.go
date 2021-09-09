package dto

type CreateLogRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
	Date string `json:"date" form:"date" binding:"required"`
}

type LogResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
}
