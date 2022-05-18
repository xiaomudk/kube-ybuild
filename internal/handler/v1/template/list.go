package template

import (
	"github.com/labstack/echo/v4"
	"github.com/xiaomudk/kube-ybuild/internal/service"
	"github.com/xiaomudk/kube-ybuild/pkg/errcode"
	"github.com/xiaomudk/kube-ybuild/pkg/logs"
)

func List(c echo.Context) error {
	logs.Info("List function called.")

	tpl, err := service.Svc.Templates().ListTemplate(c.Request().Context())
	if err != nil {
		response.Error(c, errcode.ErrInternalServer.WithDetails(err.Error()))
		return nil
	}
	response.Success(c, tpl)
	return nil
}
