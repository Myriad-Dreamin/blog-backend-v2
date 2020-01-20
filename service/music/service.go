//go:generate package-attach-to-path -generate_register_map
package musicservice

import (
	"github.com/Myriad-Dreamin/minimum-lib/module"
	"github.com/Myriad-Dreamin/blog-backend-v2/config"
	"github.com/Myriad-Dreamin/blog-backend-v2/control"
	base_service "github.com/Myriad-Dreamin/blog-backend-v2/lib/base-service"
	"github.com/Myriad-Dreamin/blog-backend-v2/model"
	"github.com/Myriad-Dreamin/blog-backend-v2/types"
)

type Service struct {
	base_service.CRUDService
	base_service.ListService
	db     *model.MusicDB
	cfg    *config.ServerConfig
	logger types.Logger
	key    string
}

func (svc *Service) MusicServiceSignatureXXX() interface{} { return svc }

func NewService(m module.Module) (control.MusicService, error) {
	var a = new(Service)
	provider := m.Require(config.ModulePath.Provider.Model).(*model.Provider)
	a.logger = m.Require(config.ModulePath.Global.Logger).(types.Logger)
	a.cfg = m.Require(config.ModulePath.Global.Configuration).(*config.ServerConfig)
	a.key = "mid"
	a.db = provider.MusicDB()
	a.CRUDService = base_service.NewCRUDService(a, a.key)
	a.ListService = base_service.NewListService(a, a.db.FilterI)
	return a, nil
}
