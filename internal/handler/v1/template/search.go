package template

import (
	"github.com/labstack/echo/v4"
	"github.com/xiaomudk/kube-ybuild/internal/service"
	"github.com/xiaomudk/kube-ybuild/pkg/errcode"
	"github.com/xiaomudk/kube-ybuild/pkg/logs"
)

func Search(c echo.Context) error {
	var req SearchRequest
	if err := c.Bind(&req); err != nil {
		logs.Warnf("search template bind param err: %v", err)
		response.Error(c, errcode.ErrInvalidParam)
		return nil
	}

	if err := c.Validate(&req); err != nil {
		logs.Warnf("search template validate param err: %v", err)
		response.Error(c, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return nil
	}
	logs.Infof("search template req: %#v", req)

	tpl, err := service.Svc.Templates().SearchTemplateByName(c.Request().Context(), req.TemplateName)
	if err != nil {
		response.Error(c, errcode.ErrInternalServer.WithDetails(err.Error()))
		return nil
	}

	response.Success(c, tpl)
	return nil
}
