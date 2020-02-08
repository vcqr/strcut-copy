package dto

import (
	"time"
)

type ArticleDto struct {
	Id        int64             `json:"id"`
	Title     string            `json:"title"`
	CatId     int64             `json:"cat_id"`
	Content   string            `json:"content"`
	Hits      int64             `json:"hits"`
	Author    []*User           `json:"author"`
	Coords    [2]*Point         `json:"coords"`
	Comment   map[int64]Comment `json:"comment"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type Comment struct {
	Author *User
	Text   string
}

type User struct {
	Id       int64  `json:"id"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
}

type Point struct {
	X string
	Y string
}
