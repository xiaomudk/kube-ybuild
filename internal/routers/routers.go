package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xiaomudk/kube-ybuild/internal/handler/v1/template"

	"github.com/xiaomudk/kube-ybuild/internal/handler/v1/user"
	"github.com/xiaomudk/kube-ybuild/pkg/app"
)

// Init initialize the routing of this application.
func Init(e *echo.Echo) {
	// HealthCheck 健康检查路由
	e.GET("/health", app.HealthCheck)
	e.Any("/*", app.RouteNotFound)
	apiV1 := e.Group("/api/v1")

	// 认证相关路由
	apiV1.POST("/register", user.Register)
	apiV1.POST("/login", user.Login)

	apiV1User := apiV1.Group("/users")
	apiV1Tpl := apiV1.Group("/templates")

	// 用户
	apiV1User.Use(middleware.JWT([]byte("secret")))

	apiV1User.GET("/info", user.Info)
	apiV1User.GET("/:id", user.Get)
	apiV1User.PUT("/:id", user.Update)
	apiV1User.DELETE("/:id", user.Delete)

	// 模板
	apiV1Tpl.POST("/create", template.Create)
	apiV1Tpl.PUT("/:id", template.Update)
	apiV1Tpl.GET("/:id", template.Get)
	apiV1Tpl.DELETE("/:id", template.Delete)
	apiV1Tpl.GET("/list", template.List)
	apiV1Tpl.GET("/search", template.Search)
}
