package musicservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/blog-backend-v2/control"
	base_service "github.com/Myriad-Dreamin/blog-backend-v2/lib/base-service"
	"github.com/Myriad-Dreamin/blog-backend-v2/model"
	"github.com/Myriad-Dreamin/blog-backend-v2/types"
)

func (svc *Service) CreateEntity(id uint) base_service.CRUDEntity {
	return &model.Music{ID: id}
}

func (svc *Service) GetEntity(id uint) (base_service.CRUDEntity, error) {
	return svc.db.ID(id)
}

func (svc *Service) ResponsePost(obj base_service.CRUDEntity) interface{} {
	return svc.AfterPost(control.SerializePostMusicReply(types.CodeOK, obj.(*model.Music)))
}

func (svc *Service) DeleteHook(c controller.MContext, obj base_service.CRUDEntity) bool {
	return svc.deleteHook(c, obj.(*model.Music))
}

func (svc *Service) ResponseGet(_ controller.MContext, obj base_service.CRUDEntity) interface{} {
	return control.SerializeGetMusicReply(types.CodeOK, obj.(*model.Music))
}

func (svc *Service) ResponseInspect(_ controller.MContext, obj base_service.CRUDEntity) interface{} {
	return control.SerializeInspectMusicReply(types.CodeOK, obj.(*model.Music))
}

func (svc *Service) ProcessListResults(_ controller.MContext, result interface{}) interface{} {
	return control.PSerializeListMusicsReply(types.CodeOK, result.([]model.Music))
}

func (svc *Service) CreateFilter() interface{} {
	return new(model.Filter)
}

func (svc *Service) GetPutRequest() interface{} {
	return new(control.PutMusicRequest)
}

func (svc *Service) FillPutFields(c controller.MContext, music base_service.CRUDEntity, req interface{}) []string {
	return svc.fillPutFields(c, music.(*model.Music), req.(*control.PutMusicRequest))
}

func (svc *Service) ListMusics(c controller.MContext) {
	svc.List(c)
	return
}

func (svc *Service) GetMusic(c controller.MContext) {
	svc.Get(c)
	return
}

func (svc *Service) PutMusic(c controller.MContext) {
	svc.Put(c)
	return
}

func (svc *Service) PostMusic(c controller.MContext) {
	svc.Post(c)
	return
}

func (svc *Service) InspectMusic(c controller.MContext) {
	svc.Inspect(c)
	return
}
