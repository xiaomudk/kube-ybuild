package user

import (
	"github.com/labstack/echo/v4"

	"github.com/xiaomudk/kube-ybuild/internal/ecode"
	"github.com/xiaomudk/kube-ybuild/internal/service"
	"github.com/xiaomudk/kube-ybuild/pkg/errcode"
	"github.com/xiaomudk/kube-ybuild/pkg/logs"
)

// Register 注册
// @Summary 注册
// @Description 用户注册
// @Tags 用户
// @Produce  json
// @Param req body RegisterRequest true "请求参数"
// @Success 200 {object} model.UserInfo "用户信息"
// @Router /Register [post]
func Register(c echo.Context) error {
	// Binding the data with the u struct.
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		logs.Warnf("register bind param err: %v", err)
		response.Error(c, errcode.ErrInvalidParam)
		return nil
	}

	if err := c.Validate(&req); err != nil {
		logs.Warnf("register validate param err: %v", err)
		response.Error(c, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return nil
	}

	logs.Infof("register req: %#v", req)

	// 两次密码是否正确
	if req.Password != req.ConfirmPassword {
		logs.Warnf("twice password is not same")
		response.Error(c, ecode.ErrTwicePasswordNotMatch)
		return nil
	}

	err := service.Svc.Users().Register(c.Request().Context(), req.Username, req.Email, req.Password)
	if err != nil {
		logs.Warnf("register err: %v", err)
		response.Error(c, ecode.ErrRegisterFailed.WithDetails(err.Error()))
		return nil
	}

	response.Success(c, nil)
	return nil
}
