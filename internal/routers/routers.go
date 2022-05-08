package routers

import (
	"github.com/labstack/echo/v4"

	"github.com/xiaomudk/kube-ybuild/pkg/app"
)

// Init initialize the routing of this application.
func Init(e *echo.Echo) {
	// HealthCheck 健康检查路由
	e.GET("/health", app.HealthCheck)
	e.Any("/*", app.RouteNotFound)
}
