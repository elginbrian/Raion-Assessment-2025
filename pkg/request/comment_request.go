package request

type CreateCommentRequest struct {
	PostID  string `json:"post_id" validate:"required"`
	Content string `json:"content" validate:"required"`
}
