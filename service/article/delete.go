package articleservice

import (
	"github.com/Myriad-Dreamin/blog-backend-v2/lib/serial"
	"github.com/Myriad-Dreamin/blog-backend-v2/model"
	"github.com/Myriad-Dreamin/blog-backend-v2/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
)

func (svc *Service) deleteHook(c controller.MContext, article *model.Article) (canDelete bool) {
	c.AbortWithStatusJSON(http.StatusOK, serial.ErrorSerializer{
		Code:  types.CodeDeleteError,
		Error: "generated delete api has not been implemented yet",
	})
	return false
}
