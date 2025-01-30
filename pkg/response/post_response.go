package response

import "time"

type GetAllPostsResponse struct {
	Status string `json:"status"`
	Data   []Post `json:"data"`
	Code   int    `json:"code"`
}

type SearchPostsResponse struct {
	Status string `json:"status"`
	Data   []Post `json:"data"`
	Code   int    `json:"code"` 
}

type GetPostByIDResponse struct {
	Status string `json:"status"`
	Data   Post   `json:"data"`
	Code   int    `json:"code"` 
}

type CreatePostResponse struct {
	Status string `json:"status"`
	Data   Post   `json:"data"`
	Code   int    `json:"code"` 
}

type UpdatePostResponse struct {
	Status string `json:"status"`
	Data   Post   `json:"data"`
	Code   int    `json:"code"`
}

type DeletePostResponse struct {
	Status string       `json:"status"`
	Data   DeletePostData `json:"data"`
	Code   int          `json:"code"` 
}

type DeletePostData struct {
	Message string `json:"message"`
	Code    int    `json:"code"` 
}

type Post struct {
	ID        string       `json:"id"`
	UserID    string       `json:"user_id"`
	Caption   string       `json:"caption"`
	ImageURL  string       `json:"image_url"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}