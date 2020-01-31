package articleservice

import (
	"github.com/Myriad-Dreamin/blog-backend-v2/lib/serial"
	ginhelper "github.com/Myriad-Dreamin/blog-backend-v2/service/gin-helper"
	"github.com/Myriad-Dreamin/blog-backend-v2/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

func (svc *Service) PutArticleContent(c controller.MContext) {
	id, ok := ginhelper.ParseUint(c, svc.key)
	if !ok {
		return
	}

	obj, err := svc.db.ID(id)
	if ginhelper.MaybeSelectError(c, obj, err) {
		return
	}

	file, err := c.FormFile("upload")
	if err != nil {
		c.JSON(http.StatusOK, serial.ErrorSerializer{
			Code:  types.CodeUploadFileError,
			Error: err.Error(),
		})
		return
	}

	if err = c.SaveUploadedFile(file,
		filepath.Join(
			svc.cfg.BaseParametersConfig.ArticlesPath,
			strconv.Itoa(int(id))+".md")); err != nil {
		c.JSON(http.StatusOK, serial.ErrorSerializer{
			Code:  types.CodeFSExecError,
			Error: err.Error(),
		})
		return
	}

	obj.UpdatedAt = time.Now()
	if !ginhelper.UpdateFields(c, obj, []string{"updated_at"}) {
		return
	}

	c.JSON(http.StatusOK, &ginhelper.ResponseOK)
}
