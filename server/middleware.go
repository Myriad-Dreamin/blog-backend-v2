package server

import (
	"github.com/Myriad-Dreamin/blog-backend-v2/config"
	"github.com/Myriad-Dreamin/blog-backend-v2/lib/jwt"
	ginhelper "github.com/Myriad-Dreamin/blog-backend-v2/service/gin-helper"
	"github.com/Myriad-Dreamin/blog-backend-v2/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/gin-contrib/cors"
	//"github.com/Myriad-Dreamin/gin-middleware/auth/privileger"
	"strconv"
)

func (srv *Server) PrepareMiddleware() bool {
	srv.jwtMW = jwt.NewMiddleWare(func() *jwt.CustomClaims {
		var cc = new(jwt.CustomClaims)
		cc.CustomField = &types.CustomFields{}
		return cc
	}, func(c controller.MContext, cc *jwt.CustomClaims) error {
		c.Set("uid", strconv.FormatInt(cc.CustomField.(*types.CustomFields).UID, 10))
		return nil
	})
	srv.jwtMW.ExpireSecond = 3600 * 24 * 7
	srv.jwtMW.RefreshSecond = 3600 * 24 * 7

	srv.routerAuthMW = controller.NewMiddleware(srv.ModelProvider.Enforcer(),
		"user:", "uid", ginhelper.MissID, ginhelper.AuthFailed)

	srv.corsMW = cors.New(cors.Config{
		//AllowAllOrigins: true,
		AllowOriginFunc: func(origin string) bool { return true },
		AllowOrigins: []string{
			"http://myriaddreamin.com:80", "https://myriaddreamin.com:80",
			"http://www.myriaddreamin.com:80", "https://www.myriaddreamin.com:80"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"X-Total-Count"},
		AllowCredentials: true,
	})

	srv.Module.Provide(config.ModulePath.Middleware.JWT, srv.jwtMW)
	srv.Module.Provide(config.ModulePath.Middleware.RouteAuth, srv.routerAuthMW)
	srv.Module.Provide(config.ModulePath.Middleware.CORS, srv.corsMW)
	return true
}
