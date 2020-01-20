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

	//Title string `dorm:"title" gorm:"title;not_null"`
	//Intro string `dorm:"intro" gorm:"intro;type:text;not_null"`
	//Category string `dorm:"category" gorm:"category;not_null"`
	//FilePath string `dorm:"file_path" gorm:"file_path;not_null"`
	// fill here
	return obj
}

func (svc *Service) AfterPost(obj interface{}) interface{} {
	return obj
}
