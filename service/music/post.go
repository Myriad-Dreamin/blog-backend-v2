package musicservice

import (
	"github.com/Myriad-Dreamin/blog-backend-v2/control"
	base_service "github.com/Myriad-Dreamin/blog-backend-v2/lib/base-service"
	"github.com/Myriad-Dreamin/blog-backend-v2/model"
	ginhelper "github.com/Myriad-Dreamin/blog-backend-v2/service/gin-helper"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

func (svc *Service) SerializePost(c controller.MContext) base_service.CRUDEntity {
	var req control.PostMusicRequest
	if !ginhelper.BindRequest(c, &req) {
		return nil
	}

	var obj = new(model.Music)
	// fill here
	obj.RecommendType = req.RecommendType
	obj.Category = req.Category
	obj.Title = req.Title
	obj.Artist = req.Artist
	obj.ReferenceID = req.ReferenceID
	obj.TrackName = req.TrackName
	obj.Comment = req.Comment
	return obj
}

func (svc *Service) AfterPost(obj interface{}) interface{} {
	return obj
}
