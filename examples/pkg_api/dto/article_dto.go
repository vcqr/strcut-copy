package dto

import (
	"time"
)

type ArticleDto struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	CatId     int64     `json:"cat_id"`
	Content   string    `json:"content"`
	Hits      int64     `json:"hits"`
	Author    *User     `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	Id       int64  `json:"id"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
}
