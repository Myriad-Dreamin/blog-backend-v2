package server

import (
	"fmt"
	"github.com/Myriad-Dreamin/blog-backend-v2/service"
	"github.com/Myriad-Dreamin/functional-go"
)

type serviceResult struct {
	serviceName string
	functional.DecayResult
}

func (srv *Server) PrepareService() bool {
	for _, serviceResult := range []serviceResult{
		{"musicService", functional.Decay(service.NewMusicService(srv.Module))},
		{"articleService", functional.Decay(service.NewArticleService(srv.Module))},
		{"userService", functional.Decay(service.NewUserService(srv.Module))},
		{"authService", functional.Decay(service.NewAuthService(srv.Module))},
		{"objectService", functional.Decay(service.NewObjectService(srv.Module))},
	} {
		// build Router failed when requesting service with database, report and return
		if serviceResult.Err != nil {
			srv.Logger.Debug(fmt.Sprintf("get %T service error", serviceResult.First), "error", serviceResult.Err)
			return false
		}
		srv.ServiceProvider.Register(serviceResult.serviceName, serviceResult.First)
	}
	return true
}
