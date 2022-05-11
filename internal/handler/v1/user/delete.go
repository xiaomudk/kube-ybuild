package user

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"

	"github.com/xiaomudk/kube-ybuild/internal/service"
	"github.com/xiaomudk/kube-ybuild/pkg/errcode"
	"github.com/xiaomudk/kube-ybuild/pkg/logs"
)

// Delete 更新用户信息
// @Summary Delete a user by ID
// @Description Delete a user by ID
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param id path uint64 true "The user's database id index num"
// @Success 200 {object} app.Response "{"code":0,"message":"OK","data":null}"
// @Router /users/{id} [delete]
func Delete(c echo.Context) error {
	// Get the user id from the url parameter.
	userID := cast.ToUint64(c.Param("id"))

	err := service.Svc.Users().DeleteUser(context.TODO(), userID)
	if err != nil {
		logs.Warnf("[user] update user err, %v", err)
		response.Error(c, errcode.ErrInternalServer)
		return nil
	}

	response.Success(c, userID)
	return nil
}
