package musicservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/blog-backend-v2/control"
	base_service "github.com/Myriad-Dreamin/blog-backend-v2/lib/base-service"
	"github.com/Myriad-Dreamin/blog-backend-v2/model"
	ginhelper "github.com/Myriad-Dreamin/blog-backend-v2/service/gin-helper"
)

func (svc *Service) SerializePost(c controller.MContext) base_service.CRUDEntity {
	var req control.PostMusicRequest
	if !ginhelper.BindRequest(c, &req) {
		return nil
	}

	var obj = new(model.Music)
	// fill here
	return obj
}

func (svc *Service) AfterPost(obj interface{}) interface{} {
	return obj
}
