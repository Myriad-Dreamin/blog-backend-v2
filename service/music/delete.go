package musicservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/blog-backend-v2/lib/serial"
	"github.com/Myriad-Dreamin/blog-backend-v2/model"
	"github.com/Myriad-Dreamin/blog-backend-v2/types"
	"net/http"
)

func (svc *Service) deleteHook(c controller.MContext, music *model.Music) (canDelete bool) {
	c.AbortWithStatusJSON(http.StatusOK, serial.ErrorSerializer{
		Code:  types.CodeDeleteError,
		Error: "generated delete api has not been implemented yet",
	})
	return false
}
