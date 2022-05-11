package app

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/xiaomudk/kube-ybuild/pkg/errcode"
	"github.com/xiaomudk/kube-ybuild/pkg/utils"
)

type ResponseHander echo.Context

var resp *Response

func init() {
	resp = NewResponse()
}

// Response define a response struct
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Details []string    `json:"details,omitempty"`
}

// NewResponse return a response
func NewResponse() *Response {
	return &Response{}
}

// Success return a success response
func Success(c ResponseHander, data interface{}) { resp.Success(c, data) }
func (r *Response) Success(c ResponseHander, data interface{}) {
	if data == nil {
		data = map[string]string{}
	}

	c.JSON(http.StatusOK, Response{
		Code:    errcode.Success.Code(),
		Message: errcode.Success.Msg(),
		Data:    data,
	})
}

// Error return a error response
func Error(c ResponseHander, err error) { resp.Error(c, err) }
func (r *Response) Error(c ResponseHander, err error) {
	if err == nil {
		c.JSON(http.StatusOK, Response{
			Code:    errcode.Success.Code(),
			Message: errcode.Success.Msg(),
			Data:    map[string]string{},
		})
		return
	}

	if v, ok := err.(*errcode.Error); ok {
		response := Response{
			Code:    v.Code(),
			Message: v.Msg(),
			Data:    map[string]string{},
			Details: []string{},
		}
		details := v.Details()
		if len(details) > 0 {
			response.Details = details
		}
		c.JSON(errcode.ToHTTPStatusCode(v.Code()), response)
		return
	}
}

// RouteNotFound 未找到相关路由
func RouteNotFound(c echo.Context) error {
	return c.String(http.StatusNotFound, "the route not found")
}

// healthCheckResponse 健康检查响应结构体
type healthCheckResponse struct {
	Status   string `json:"status"`
	Hostname string `json:"hostname"`
}

// HealthCheck will return OK if the underlying BoltDB is healthy.
// At least healthy enough for demoing purposes.
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, healthCheckResponse{Status: "UP", Hostname: utils.GetHostname()})
}
