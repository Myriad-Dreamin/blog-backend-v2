package musicservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/blog-backend-v2/control"
	"github.com/Myriad-Dreamin/blog-backend-v2/model"
)

func (svc *Service) fillPutFields(
	c controller.MContext, music *model.Music,
	req *control.PutMusicRequest) (fields []string) {
	return nil
}
