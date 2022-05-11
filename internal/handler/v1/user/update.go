package user

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"

	"github.com/xiaomudk/kube-ybuild/internal/service"
	"github.com/xiaomudk/kube-ybuild/pkg/errcode"
	"github.com/xiaomudk/kube-ybuild/pkg/logs"
)

// Update 更新用户信息
// @Summary Update a user info by the user identifier
// @Description Update a user by ID
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param id path uint64 true "The user's database id index num"
// @Param user body model.UserBaseModel true "The user info"
// @Success 200 {object} app.Response "{"code":0,"message":"OK","data":null}"
// @Router /users/{id} [put]
func Update(c echo.Context) error {
	// Get the user id from the url parameter.
	userID := cast.ToUint64(c.Param("id"))

	// Binding the user data.
	var req UpdateRequest
	if err := c.Bind(&req); err != nil {
		logs.Warnf("bind request param err: %+v", err)
		response.Error(c, errcode.ErrInvalidParam)
		return nil
	}
	logs.Infof("user update req: %#v", req)

	userMap := make(map[string]interface{})
	userMap["phone"] = req.Phone
	userMap["email"] = req.Email
	userMap["sex"] = req.Sex
	err := service.Svc.Users().UpdateUser(context.TODO(), userID, userMap)
	if err != nil {
		logs.Warnf("[user] update user err, %v", err)
		response.Error(c, errcode.ErrInternalServer)
		return nil
	}

	response.Success(c, userID)
	return nil
}
