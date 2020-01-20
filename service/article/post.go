package articleservice

import (
	"github.com/Myriad-Dreamin/blog-backend-v2/control"
	base_service "github.com/Myriad-Dreamin/blog-backend-v2/lib/base-service"
	"github.com/Myriad-Dreamin/blog-backend-v2/model"
	ginhelper "github.com/Myriad-Dreamin/blog-backend-v2/service/gin-helper"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

func (svc *Service) SerializePost(c controller.MContext) base_service.CRUDEntity {
	var req control.PostArticleRequest
	if !ginhelper.BindRequest(c, &req) {
		return nil
	}

	var obj = new(model.Article)
	obj.Title = req.Title
	obj.Intro = req.Intro
	obj.Category = req.Category
	obj.FilePath = req.FilePath
	// fill here
	return obj
}

func (svc *Service) AfterPost(obj interface{}) interface{} {
	return obj
}
