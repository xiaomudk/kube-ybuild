package template

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"github.com/xiaomudk/kube-ybuild/internal/ecode"
	"github.com/xiaomudk/kube-ybuild/internal/repository"
	"github.com/xiaomudk/kube-ybuild/internal/service"
	"github.com/xiaomudk/kube-ybuild/pkg/errcode"
	"github.com/xiaomudk/kube-ybuild/pkg/logs"
)

func Get(c echo.Context) error {
	logs.Info("Get function called.")

	tplID := cast.ToUint64(c.Param("id"))
	if tplID == 0 {
		response.Error(c, errcode.ErrInvalidParam)
		return nil
	}

	// Get the template by the `tpl_id` from the database.
	tpl, err := service.Svc.Templates().GetTemplateById(c.Request().Context(), tplID)
	if errors.Is(err, repository.ErrNotFound) {
		logs.Errorf("get template info err: %+v", err)
		response.Error(c, ecode.ErrTemplateNotFound)
		return nil
	}
	if err != nil {
		response.Error(c, errcode.ErrInternalServer.WithDetails(err.Error()))
		return nil
	}

	response.Success(c, tpl)
	return nil
}
