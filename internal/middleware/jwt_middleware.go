package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/xiaomudk/kube-ybuild/internal/service"
	"github.com/xiaomudk/kube-ybuild/pkg/utils"
	"net/http"
)

func TokenValidateMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("token")
		if token == "" {
			return c.JSON(http.StatusForbidden, map[string]interface{}{
				"code":    403,
				"message": "header not take token,access denied",
				"data":    "{}",
			})
		}

		user, err := utils.ParseToken(token)

		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]interface{}{
				"code":    403,
				"message": err.Error(),
				"data":    "{}",
			})
		}
		u, err := service.Svc.Users().GetUserByID(c.Request().Context(), user.UserId)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": err.Error(),
				"data":    "{}",
			})
		}
		fmt.Println(u)
		return next(c)
	}
}
