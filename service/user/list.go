package userservice

import (
	"github.com/Myriad-Dreamin/blog-backend-v2/control"
	"github.com/Myriad-Dreamin/blog-backend-v2/model"
	"github.com/Myriad-Dreamin/blog-backend-v2/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

func (srv *Service) ProcessListResults(c controller.MContext, result interface{}) interface{} {
	return control.SerializeListUsersReply(
		types.CodeOK,
		control.PackSerializeListUserReply(result.([]model.User)))
}
