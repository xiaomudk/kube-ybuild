package user

import (
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"

	"github.com/xiaomudk/kube-ybuild/internal/ecode"
	"github.com/xiaomudk/kube-ybuild/internal/repository"
	"github.com/xiaomudk/kube-ybuild/internal/service"
	"github.com/xiaomudk/kube-ybuild/pkg/errcode"
	"github.com/xiaomudk/kube-ybuild/pkg/logs"
)

// Info 获取登录用户信息
// @Summary 通过用户token获取用户信息
// @Description Get an user by user token
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param id path string true "用户id"
// @Success 200 {object} model.UserInfo "用户信息"
// @Router /users/info [get]
func Info(c echo.Context) error {
	logs.Info("Info function called.")

	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		logs.Warnf("failed to convert token %v", token)
		response.Error(c, errcode.ErrInvalidParam)
		return nil
	}

	// Get the user by the `user_id` from the database.
	u, err := service.Svc.Users().GetUserByToken(c.Request().Context(), token)
	if errors.Is(err, repository.ErrNotFound) {
		logs.Errorf("get user info err: %+v", err)
		response.Error(c, ecode.ErrUserNotFound)
		return nil
	}
	if err != nil {
		response.Error(c, errcode.ErrInternalServer.WithDetails(err.Error()))
		return nil
	}

	response.Success(c, u)
	return nil
}
