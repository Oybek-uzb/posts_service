package model

type Post struct {
	Id     int32  `json:"id"`
	UserId int32  `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
