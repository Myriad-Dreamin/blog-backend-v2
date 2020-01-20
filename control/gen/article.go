package main

import (
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/blog-backend-v2/model"
)

type ArticleCategories struct {
	artisan.VirtualService
	List    artisan.Category
	Post    artisan.Category
	Inspect artisan.Category
	IdGroup artisan.Category
}

func DescribeArticleService(base string) artisan.ProposingService {
	var articleModel = new(model.Article)
	svc := &ArticleCategories{
		List: artisan.Ink().
			Path("article-list").
			Method(artisan.POST, "ListArticles",
				artisan.QT("ListArticlesRequest", model.Filter{}),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("articles", articleModel)),
				),
			),
		Post: artisan.Ink().
			Path("article").
			Method(artisan.POST, "PostArticle",
				artisan.Request(),
				artisan.Reply(
					codeField,
					artisan.Param("article", &articleModel),
				),
			),
		Inspect: artisan.Ink().Path("article/:aid/inspect").
			Method(artisan.GET, "InspectArticle",
				artisan.Reply(
					codeField,
					artisan.Param("article", &articleModel),
				),
			),
		IdGroup: artisan.Ink().
			Path("article/:aid").
			Method(artisan.GET, "GetArticle",
				artisan.Reply(
					codeField,
					artisan.Param("article", &articleModel),
				)).
			Method(artisan.PUT, "PutArticle",
				artisan.Request()).
			Method(artisan.DELETE, "Delete"),
	}
	svc.Name("ArticleService").Base(base) //.
	//UseModel(serial.Model(serial.Name("article"), &articleModel))
	return svc
}
