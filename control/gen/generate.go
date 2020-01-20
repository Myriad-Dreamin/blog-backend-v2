package main

import (
	"fmt"
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/blog-backend-v2/types"
)

var codeField = artisan.Param("code", new(types.CodeRawType))
var required = artisan.Tag("binding", "required")

func main() {
	v1 := "v1"

	//instantiate
    musicCate := DescribeMusicService(v1)
    articleCate := DescribeArticleService(v1)
	userCate := DescribeUserService(v1)
	objectCate := DescribeObjectService(v1)

	//to files
    musicCate.ToFile("music.go")
    articleCate.ToFile("article.go")
	userCate.ToFile("user.go")
	objectCate.ToFile("object.go")

	err := artisan.NewService(
        musicCate,
        articleCate,
		userCate,
		objectCate,
	).Publish()
	if err != nil {
		fmt.Println(err)
	}
}
