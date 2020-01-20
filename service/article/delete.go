package articleservice

import (
	"github.com/Myriad-Dreamin/blog-backend-v2/model"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

func (svc *Service) deleteHook(c controller.MContext, article *model.Article) (canDelete bool) {

	return true
}
