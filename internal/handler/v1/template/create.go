package template

import (
	"github.com/labstack/echo/v4"
	"github.com/xiaomudk/kube-ybuild/internal/ecode"
	"github.com/xiaomudk/kube-ybuild/internal/service"
	"github.com/xiaomudk/kube-ybuild/pkg/errcode"
	"github.com/xiaomudk/kube-ybuild/pkg/logs"
)

func Create(c echo.Context) error {
	var req CreateRequest
	if err := c.Bind(&req); err != nil {
		logs.Warnf("create template bind param err: %v", err)
		response.Error(c, errcode.ErrInvalidParam)
		return nil
	}

	if err := c.Validate(&req); err != nil {
		logs.Warnf("create template validate param err: %v", err)
		response.Error(c, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return nil
	}

	logs.Infof("create template req: %#v", req)
	err := service.Svc.Templates().CreateTemplate(c.Request().Context(), req.TemplateName, req.TemplateContent)
	if err != nil {
		logs.Warnf("create template: %v", err)
		response.Error(c, ecode.ErrTemplateCreate.WithDetails(err.Error()))
		return nil
	}
	response.Success(c, nil)
	return nil
}
