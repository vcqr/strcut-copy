package pkg_article

import (
	"time"

	"structutils/examples/pkg_user"
)

type Article struct {
	Id        int64                       `json:"id"`
	Title     string                      `json:"title"`
	CatId     int64                       `json:"cat_id"`
	Content   string                      `json:"content"`
	Hits      int64                       `json:"hits"`
	Author    []*pkg_user.User            `json:"author"`
	Coords    [2]*Point                   `json:"coords"`
	CreatedAt time.Time                   `json:"created_at"`
	UpdatedAt time.Time                   `json:"updated_at"`
	Comment   map[int64]*pkg_user.Comment `json:"comment"`
}

type Point struct {
	X string
	Y string
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
		Author: []*pkg_user.User{
			&pkg_user.User{
				Id:       1,
				NickName: "Tom",
				Avatar:   "tom.png",
				Email:    "demo@test.com",
			},
			&pkg_user.User{
				Id:       2,
				NickName: "vc",
				Avatar:   "vc.png",
				Email:    "vc@test.com",
			},
		},
		CreatedAt: now,
		UpdatedAt: now,
		Comment: map[int64]*pkg_user.Comment{
			1000: &pkg_user.Comment{
				Author: &pkg_user.User{
					Id:       1,
					NickName: "游客1",
					Avatar:   "d.png",
					Email:    "xxx@xxx.com",
				},
				Text: "他很懒，什么也没有留下",
			},
			1001: &pkg_user.Comment{
				Author: &pkg_user.User{
					Id:       2,
					NickName: "游客2",
					Avatar:   "d.png",
					Email:    "xxx@xxx.com",
				},
				Text: "空空如也",
			},
		},
		Coords: [2]*Point{
			&Point{
				X: "116.404",
				Y: "39.915",
			},
			&Point{
				X: "117.404",
				Y: "40.915",
			},
		},
	}
}
