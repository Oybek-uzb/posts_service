package model

type Post struct {
	Id     int32  `json:"id"`
	UserId int32  `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Body struct {
	Meta `json:"meta"`
	Data []Post `json:"data"`
}

type Meta struct {
	Pagination `json:"pagination"`
}

type Pagination struct {
	Total int32 `json:"total"`
	Pages int32 `json:"pages"`
	Page  int32 `json:"page"`
	Limit int32 `json:"limit"`
	Links `json:"links"`
}

type Links struct {
	Previous string `json:"previous"`
	Current  string `json:"current"`
	Next     string `json:"next"`
}
