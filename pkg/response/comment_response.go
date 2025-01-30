package response

import "time"

type GetCommentsResponse struct {
	Status string 		`json:"status"`
	Data   []Comment 	`json:"data"`
}

type CreateCommentResponse struct {
	Status string 	`json:"status"`
	Data   Comment 	`json:"data"`
}

type DeleteCommentResponse struct {
	Status string `json:"status"`
    Message string `json:"message"`
}

type Comment struct {
    ID        string    `json:"id"`
    UserID    string    `json:"user_id"`
    PostID    string    `json:"post_id"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}