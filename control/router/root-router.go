package router

import (
	"github.com/Myriad-Dreamin/blog-backend-v2/config"
	"github.com/Myriad-Dreamin/blog-backend-v2/lib/jwt"
	"github.com/Myriad-Dreamin/blog-backend-v2/service"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseH struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware
}

func (r *BaseH) GetRouter() *Router {
	return r.Router
}

func (r *BaseH) GetAuthRouter() *Router {
	return r.AuthRouter
}

func (r *BaseH) GetAuth() *Middleware {
	return r.Auth
}

type RootRouter struct {
	H
	Root *Router

	//ObjectRouter *ObjectRouter
	UserRouter    *UserRouter
	AuthRouter    *AuthRouter
	ArticleRouter *ArticleRouter
	MusicRouter   *MusicRouter

	Ping     *LeafRouter
	Images   *LeafRouter
	Musics   *LeafRouter
	Articles *LeafRouter
}

// @title Ping
// @description result
func PingFunc(c controller.MContext) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func NewRootRouter(m module.Module) (r *RootRouter) {
	rr := controller.NewRouterGroup()
	apiRouterV1 := rr.Group("/v1")
	b := m.Require(config.ModulePath.Middleware.JWT).(*jwt.Middleware).Build()
	authRouterV1 := apiRouterV1.Group("", b)

	r = &RootRouter{
		Root: rr,
		H: &BaseH{
			Router:     apiRouterV1,
			AuthRouter: authRouterV1,
			Auth:       m.Require(config.ModulePath.Middleware.RouteAuth).(*Middleware),
		},
	}

	r.Ping = r.Root.GET("/ping", PingFunc)

	//r.ObjectRouter = BuildObjectRouter(r, serviceProvider)

	serviceProvider := m.Require(config.ModulePath.Provider.Service).(*service.Provider)

	r.UserRouter = BuildUserRouter(r, serviceProvider)
	r.ArticleRouter = BuildArticleRouter(r, serviceProvider)
	r.MusicRouter = BuildMusicRouter(r, serviceProvider)

	cfg := m.Require(config.ModulePath.Global.Configuration).(*config.ServerConfig)

	r.Images = r.GetRouter().StaticFS("images", http.Dir(cfg.BaseParametersConfig.ImagesPath))
	r.Musics = r.GetRouter().StaticFS("musics", http.Dir(cfg.BaseParametersConfig.MusicsPath))
	r.Articles = r.GetRouter().StaticFS("articles", http.Dir(cfg.BaseParametersConfig.ArticlesPath))

	ApplyAuth(r)
	return
}
