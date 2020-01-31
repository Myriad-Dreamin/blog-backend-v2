package musicservice

import (
	"github.com/Myriad-Dreamin/blog-backend-v2/model"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

func (svc *Service) deleteHook(c controller.MContext, music *model.Music) (canDelete bool) {

	return true
}
