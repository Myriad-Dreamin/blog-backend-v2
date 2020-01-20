package router

import (
	"github.com/Myriad-Dreamin/blog-backend-v2/service"
)

type ArticleRouter struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware
	IDRouter   *ArticleIDRouter

	Post    *LeafRouter
	GetList *LeafRouter
}

type ArticleIDRouter struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware

	Get    *LeafRouter
	Put    *LeafRouter
	Delete *LeafRouter
}

func BuildArticleRouter(parent H, serviceProvider *service.Provider) (router *ArticleRouter) {
	articleService := serviceProvider.ArticleService()
	router = &ArticleRouter{
		Router:     parent.GetRouter().Extend("article"),
		AuthRouter: parent.GetAuthRouter().Extend("article"),
		Auth:       parent.GetAuth().Copy(),
	}
	router.GetList = router.GET("article-list", articleService.ListArticles)
	router.Post = router.AuthRouter.POST("/article", articleService.PostArticle)

	router.IDRouter = router.IDRouter.subBuild(router, serviceProvider)

	return
}

func (*ArticleIDRouter) subBuild(parent *ArticleRouter, serviceProvider *service.Provider) (
	router *ArticleIDRouter) {

	articleService := serviceProvider.ArticleService()

	router = &ArticleIDRouter{
		Router:     parent.Group("/article/:aid"),
		AuthRouter: parent.AuthRouter.Group("/article/:aid"),
		Auth:       parent.Auth.MustGroup("article", "aid"),
	}

	router.Get = router.GET("", articleService.GetArticle)
	router.Put = router.AuthRouter.PUT("", articleService.PutArticle)
	router.Delete = router.AuthRouter.DELETE("", articleService.Delete)
	return
}

func (s *Provider) ArticleRouter() *ArticleRouter {
	return s.articleRouter
}
