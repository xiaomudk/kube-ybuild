package template

import (
	"github.com/xiaomudk/kube-ybuild/pkg/app"
)

var response = app.NewResponse()

// CreateRequest 创建用户请求
type CreateRequest struct {
	TemplateName    string `json:"templateName" form:"templateName" binding:"required"`
	TemplateContent string `json:"templateContent" form:"templateContent" binding:"required"`
}

type UpdateRequest struct {
	TemplateName    string `json:"templateName" form:"templateName"`
	TemplateContent string `json:"templateContent" form:"templateContent"`
}

type SearchRequest struct {
	TemplateName string `json:"templateName" query:"templateName"`
}
