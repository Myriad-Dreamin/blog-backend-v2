//go:generate package-attach-to-path -generate_register_map
package authservice

import (
	"github.com/Myriad-Dreamin/blog-backend-v2/config"
	"github.com/Myriad-Dreamin/blog-backend-v2/lib/jwt"
	"github.com/Myriad-Dreamin/blog-backend-v2/model"
	"github.com/Myriad-Dreamin/blog-backend-v2/model/sp-layer"
	"github.com/Myriad-Dreamin/blog-backend-v2/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"net/http"
)

type Service struct {
	cfg        *config.ServerConfig
	logger     types.Logger
	middleware *jwt.Middleware
	enforcer   *splayer.Enforcer
}

func (svc *Service) AuthSignatureXXX() interface{} { return svc }

type RefreshTokenReply struct {
	Code  types.CodeRawType `json:"code"`
	Token string            `json:"token"`
}

func (svc *Service) RefreshToken(c controller.MContext) {
	newToken, err := svc.middleware.RefreshToken(c)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, RefreshTokenReply{
		Code:  types.CodeOK,
		Token: newToken,
	})
}

func NewService(m module.Module) (a *Service, err error) {
	a = new(Service)
	a.logger = m.Require(config.ModulePath.Global.Logger).(types.Logger)
	a.cfg = m.Require(config.ModulePath.Global.Configuration).(*config.ServerConfig)
	a.enforcer = m.Require(config.ModulePath.Provider.Model).(*model.Provider).Enforcer()
	a.middleware = m.Require(config.ModulePath.Middleware.JWT).(*jwt.Middleware)
	return
}
