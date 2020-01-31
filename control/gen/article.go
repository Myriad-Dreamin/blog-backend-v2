package main

import (
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/blog-backend-v2/model"
)

type ArticleCategories struct {
	artisan.VirtualService
	List       artisan.Category
	Post       artisan.Category
	PutContent artisan.Category
	GetContent artisan.Category
	Inspect    artisan.Category
	IdGroup    artisan.Category
}

func DescribeArticleService(base string) artisan.ProposingService {
	var articleModel = new(model.Article)
	var _articleModel = new(model.Article)
	svc := &ArticleCategories{
		List: artisan.Ink().
			Path("article-list").
			Method(artisan.POST, "ListArticles",
				artisan.QT("ListArticlesRequest", model.Filter{}),
				artisan.Reply(
					codeField,
					artisan.ArrayParam(artisan.Param("articles", _articleModel)),
				),
			),
		Post: artisan.Ink().
			Path("article").
			Method(artisan.POST, "PostArticle",
				artisan.Request(
					artisan.SPsC(&articleModel.Title, &articleModel.Intro,
						&articleModel.Category, &articleModel.PublishedAt),
				),
				artisan.Reply(
					codeField,
					artisan.Param("article", &articleModel),
				),
			),
		PutContent: artisan.Ink().
			Path("article/:aid/content").
			Method(artisan.PUT, "PutArticleContent"),
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
	svc.Name("ArticleService").Base(base).
		UseModel(artisan.Model(artisan.Name("article"), &articleModel))
	return svc
}
