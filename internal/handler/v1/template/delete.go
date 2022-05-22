package template

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"github.com/xiaomudk/kube-ybuild/internal/ecode"
	"github.com/xiaomudk/kube-ybuild/internal/service"
	"github.com/xiaomudk/kube-ybuild/pkg/logs"
)

func Delete(c echo.Context) error {
	// Get the user id from the url parameter.
	tplId := cast.ToUint64(c.Param("id"))

	err := service.Svc.Templates().DeleteTemplate(context.TODO(), tplId)
	if err != nil {
		logs.Warnf("[template] delete template err, %v", err)
		response.Error(c, ecode.ErrTemplateNotFound)
		return nil
	}

	response.Success(c, tplId)
	return nil
}
