package template

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"github.com/xiaomudk/kube-ybuild/internal/service"
	"github.com/xiaomudk/kube-ybuild/pkg/errcode"
	"github.com/xiaomudk/kube-ybuild/pkg/logs"
)

func Update(c echo.Context) error {
	// Get the template id from the url parameter.
	tplId := cast.ToUint64(c.Param("id"))
	fmt.Printf("aaaaaaaaa")
	fmt.Println(tplId)

	// Binding the user data.
	var req UpdateRequest
	if err := c.Bind(&req); err != nil {
		logs.Warnf("bind request param err: %+v", err)
		response.Error(c, errcode.ErrInvalidParam)
		return nil
	}
	logs.Infof("template update req: %#v", req)

	templateMap := make(map[string]interface{})
	templateMap["TemplateName"] = req.TemplateName
	templateMap["TemplateContent"] = req.TemplateContent
	fmt.Println(templateMap)

	err := service.Svc.Templates().UpdateTemplate(context.TODO(), tplId, templateMap)
	if err != nil {
		fmt.Printf(err.Error())
		logs.Warnf("[user] update user err, %v", err)
		response.Error(c, errcode.ErrInternalServer)
		return nil
	}

	response.Success(c, tplId)
	return nil
}
