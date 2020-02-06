package pkg_article

import (
	"time"

	"structutils/examples/pkg_user"
)

type Article struct {
	Id        int64          `json:"id"`
	Title     string         `json:"title"`
	CatId     int64          `json:"cat_id"`
	Content   string         `json:"content"`
	Hits      int64          `json:"hits"`
	Author    *pkg_user.User `json:"author"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Comment   []string       `json:"comment"`
}

func NewArticle() *Article {
	return &Article{}
}

func (a *Article) GetArticle() *Article {
	now := time.Now()

	return &Article{
		Id:      10000,
		Title:   "This is test article",
		CatId:   1,
		Content: "This is an article on struct copy, which you can use as follows",
		Hits:    100,
		Comment: []string{"nice", "good"},
		Author: &pkg_user.User{
			Id:       1,
			NickName: "Tom",
			Avatar:   "tom.png",
			Email:    "demo@test.com",
		},
		CreatedAt: now,
		UpdatedAt: now,
	}
}
