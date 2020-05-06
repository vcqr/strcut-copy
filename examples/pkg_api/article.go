package main

import (
	"encoding/json"
	"log"
	"net/http"
	
	structUtil "github.com/vcqr/struct-copy"
	"github.com/vcqr/struct-copy/examples/pkg_api/dto"
	"github.com/vcqr/struct-copy/examples/pkg_article"
)

func showHandler(w http.ResponseWriter, r *http.Request) {
	articleService := pkg_article.NewArticle()
	article := articleService.GetArticle()

	var vo dto.ArticleDto
	structUtil.NewStructUtils().CopyProperties(&vo, article)

	vo.Hits++

	resp, _ := json.Marshal(vo)

	w.Write(resp)
}

func main() {
	http.HandleFunc("/article/show", showHandler)

	log.Println("Starting article server ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
