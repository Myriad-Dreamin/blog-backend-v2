package articleservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/blog-backend-v2/control"
	base_service "github.com/Myriad-Dreamin/blog-backend-v2/lib/base-service"
	"github.com/Myriad-Dreamin/blog-backend-v2/model"
	"github.com/Myriad-Dreamin/blog-backend-v2/types"
)

func (svc *Service) CreateEntity(id uint) base_service.CRUDEntity {
	return &model.Article{ID: id}
}

func (svc *Service) GetEntity(id uint) (base_service.CRUDEntity, error) {
	return svc.db.ID(id)
}

func (svc *Service) ResponsePost(obj base_service.CRUDEntity) interface{} {
	return svc.AfterPost(control.SerializePostArticleReply(types.CodeOK, obj.(*model.Article)))
}

func (svc *Service) DeleteHook(c controller.MContext, obj base_service.CRUDEntity) bool {
	return svc.deleteHook(c, obj.(*model.Article))
}

func (svc *Service) ResponseGet(_ controller.MContext, obj base_service.CRUDEntity) interface{} {
	return control.SerializeGetArticleReply(types.CodeOK, obj.(*model.Article))
}

func (svc *Service) ResponseInspect(_ controller.MContext, obj base_service.CRUDEntity) interface{} {
	return control.SerializeInspectArticleReply(types.CodeOK, obj.(*model.Article))
}

func (svc *Service) ProcessListResults(_ controller.MContext, result interface{}) interface{} {
	return control.PSerializeListArticlesReply(types.CodeOK, result.([]model.Article))
}

func (svc *Service) CreateFilter() interface{} {
	return new(model.Filter)
}

func (svc *Service) GetPutRequest() interface{} {
	return new(control.PutArticleRequest)
}

func (svc *Service) FillPutFields(c controller.MContext, article base_service.CRUDEntity, req interface{}) []string {
	return svc.fillPutFields(c, article.(*model.Article), req.(*control.PutArticleRequest))
}

func (svc *Service) ListArticles(c controller.MContext) {
	svc.List(c)
	return
}

func (svc *Service) GetArticle(c controller.MContext) {
	svc.Get(c)
	return
}

func (svc *Service) PutArticle(c controller.MContext) {
	svc.Put(c)
	return
}

func (svc *Service) PostArticle(c controller.MContext) {
	svc.Post(c)
	return
}

func (svc *Service) InspectArticle(c controller.MContext) {
	svc.Inspect(c)
	return
}
