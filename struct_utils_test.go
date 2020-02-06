package structutils

import (
	"testing"

	"structutils/examples/pkg_api/dto"
	"structutils/examples/pkg_article"

	"github.com/stretchr/testify/assert"
)

var util = NewStructUtils()

func TestStructUtils_CopyProperties(t *testing.T) {
	a := assert.New(t)

	articleService := pkg_article.NewArticle()
	article := articleService.GetArticle()

	var articleVo dto.ArticleDto
	util.CopyProperties(&articleVo, article)

	a.Equal(int64(10000), articleVo.Id)
	a.GreaterOrEqual(int64(100), article.Hits)

	a.Equal(int64(1), articleVo.Author.Id)
	a.Equal("Tom", articleVo.Author.NickName)
}

func BenchmarkStructUtils_CopyProperties(b *testing.B) {
	articleService := pkg_article.NewArticle()
	article := articleService.GetArticle()

	var articleVo dto.ArticleDto

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		util.CopyProperties(&articleVo, article)
	}
}
