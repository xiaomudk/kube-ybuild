package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Init initialize the routing of this application.
func Init(e *echo.Echo) {
	// HealthCheck 健康检查路由
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
}
