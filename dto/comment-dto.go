package dto

type CommentCreateDTO struct {
	Body string `json:"body" form:"body" binding:"required"`
	UserID uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
	articleId uint64 `json:"article_id,omitempty" form:"article_id,omitempty"`
}

type CommentUpdateDTO struct {
	ID uint64 `json:"id" form:"id" binding:"required"`
	Body string `json:"body" form:"body" binding:"required"`
	UserID uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
	articleId uint64 `json:"article_id,omitempty" form:"article_id,omitempty"`
}
