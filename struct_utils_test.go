package struct_copy

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
	
	"github.com/vcqr/struct-copy/examples/pkg_api/dto"
	"github.com/vcqr/struct-copy/examples/pkg_article"
)

var util = NewStructUtils()

type StructEasyCaseA struct {
	A int
	B float64
	C string
}
type StructEasyCaseB struct {
	a int
	B float64
	C string
}

func TestStructUtils_CopyProperties(t *testing.T) {
	a := assert.New(t)

	articleService := pkg_article.NewArticle()
	article := articleService.GetArticle()

	var articleVo dto.ArticleDto
	util.CopyProperties(&articleVo, article)

	a.Equal(int64(10000), articleVo.Id)
	a.GreaterOrEqual(int64(100), articleVo.Hits)

	// 数组
	a.Equal("116.404", articleVo.Coords[0].X)

	// 切片
	a.Equal(2, len(articleVo.Author))
	a.Equal(int64(2), articleVo.Author[1].Id)

	// map
	a.Equal(2, len(articleVo.Comment))
	a.Equal("游客1", articleVo.Comment[1000].Author.NickName)
}

func BenchmarkStructUtils_CopyProperties_Easy(b *testing.B) {
	var src = StructEasyCaseA{}
	var dest StructEasyCaseB

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		util.CopyProperties(&dest, src)
	}
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
