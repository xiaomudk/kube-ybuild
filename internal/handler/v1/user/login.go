package user

import (
	"github.com/labstack/echo/v4"

	"github.com/xiaomudk/kube-ybuild/internal/ecode"
	"github.com/xiaomudk/kube-ybuild/internal/model"
	"github.com/xiaomudk/kube-ybuild/internal/service"
	"github.com/xiaomudk/kube-ybuild/pkg/errcode"
	"github.com/xiaomudk/kube-ybuild/pkg/logs"
)

// Login 登录接口
// @Summary 用户登录接口
// @Description 登录
// @Tags 用户
// @Produce  json
// @Param req body LoginCredentials true "phone"
// @Success 200 {object} model.UserInfo "用户信息"
// @Router /users/login [post]
func Login(c echo.Context) error {
	logs.Info("Login function called.")

	// Binding the data with the u struct.
	var req LoginCredentials
	if err := c.Bind(&req); err != nil {
		logs.Warnf("login bind param err: %v", err)
		response.Error(c, errcode.ErrInvalidParam)
		return nil
	}
	if err := c.Validate(&req); err != nil {
		logs.Warnf("login validate param err: %v", err)
		response.Error(c, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return nil
	}

	logs.Infof("req %#v", req)

	// 登录
	t, err := service.Svc.Users().Login(c.Request().Context(), req.Username, req.Password)
	if err != nil {
		response.Error(c, ecode.ErrPasswordIncorrect)
		return nil
	}

	response.Success(c, model.Token{Token: t})
	return nil
}
