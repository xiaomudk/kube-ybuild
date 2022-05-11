package user

import (
	"github.com/xiaomudk/kube-ybuild/internal/model"
	"github.com/xiaomudk/kube-ybuild/pkg/app"
)

var response = app.NewResponse()

// CreateRequest 创建用户请求
type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email" form:"email"`
}

// CreateResponse 创建用户响应
type CreateResponse struct {
	Username string `json:"username"`
}

// RegisterRequest 注册
type RegisterRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

// LoginCredentials 登录
type LoginCredentials struct {
	Username string `json:"username" form:"username" binding:"required" `
	Password string `json:"password" form:"password" binding:"required" `
}

// UpdateRequest 更新请求
type UpdateRequest struct {
	Sex   int    `json:"sex"`
	Email string `json:"email"`
	Phone int64  `json:"phone"`
}

// FollowRequest 关注请求
type FollowRequest struct {
	UserID uint64 `json:"user_id"`
}

// ListResponse 通用列表resp
type ListResponse struct {
	TotalCount uint64      `json:"total_count"`
	HasMore    int         `json:"has_more"`
	PageKey    string      `json:"page_key"`
	PageValue  int         `json:"page_value"`
	Items      interface{} `json:"items"`
}

// SwaggerListResponse 文档
type SwaggerListResponse struct {
	TotalCount uint64           `json:"totalCount"`
	UserList   []model.UserInfo `json:"userList"`
}
