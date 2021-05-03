package dto

type ArticleCreateDTO struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Slug        string `json:"slug" form:"slug"`
	Description string `json:"description" form:"description" binding:"required"`
	Image       string `json:"image" form:"image" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
}

type ArticleUpdateDTO struct {
	ID          uint64 `json:"id" form:"id"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Image       string `json:"image,omitempty" form:"image,omitempty"`
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty" `
}
