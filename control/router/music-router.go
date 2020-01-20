package router

import (
	"github.com/Myriad-Dreamin/blog-backend-v2/service"
)

type MusicRouter struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware
	IDRouter   *MusicIDRouter

	Post    *LeafRouter
	GetList *LeafRouter
}

type MusicIDRouter struct {
	*Router
	AuthRouter *Router
	Auth       *Middleware

	Get    *LeafRouter
	Put    *LeafRouter
	Delete *LeafRouter
}

func BuildMusicRouter(parent H, serviceProvider *service.Provider) (router *MusicRouter) {
	musicService := serviceProvider.MusicService()
	router = &MusicRouter{
		Router:     parent.GetRouter().Extend("music"),
		AuthRouter: parent.GetAuthRouter().Extend("music"),
		Auth:       parent.GetAuth().Copy(),
	}
	router.GetList = router.GET("music-list", musicService.ListMusics)
	router.Post = router.AuthRouter.POST("/music", musicService.PostMusic)

	router.IDRouter = router.IDRouter.subBuild(router, serviceProvider)

	return
}

func (*MusicIDRouter) subBuild(parent *MusicRouter, serviceProvider *service.Provider) (
	router *MusicIDRouter) {

	musicService := serviceProvider.MusicService()

	router = &MusicIDRouter{
		Router:     parent.Group("/music/:mid"),
		AuthRouter: parent.AuthRouter.Group("/music/:mid"),
		Auth:       parent.Auth.MustGroup("music", "mid"),
	}

	router.Get = router.GET("", musicService.GetMusic)
	router.Put = router.AuthRouter.PUT("", musicService.PutMusic)
	router.Delete = router.AuthRouter.DELETE("", musicService.Delete)
	return
}

func (s *Provider) MusicRouter() *MusicRouter {
	return s.musicRouter
}
