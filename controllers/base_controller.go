package controllers

import (
	"github.com/gin-gonic/gin"
)

type Props struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type BaseController interface {
	Data(c *gin.Context, code int, data interface{}, message string)
	Error(c *gin.Context, code int, friendlyErrorForClient error)
}

type baseController struct{}

func NewBaseController() BaseController {
	return &baseController{}
}

func (bc *baseController) Data(c *gin.Context, code int, data interface{}, message string) {
	props := &Props{
		Code:    code,
		Data:    data,
		Message: message,
	}

	c.JSON(code, props)
}

func (bc *baseController) Error(c *gin.Context, code int, friendlyErrorForClient error) {
	props := &Props{
		Code:    code,
		Data:    nil,
		Message: friendlyErrorForClient.Error(),
	}

	c.AbortWithStatusJSON(code, props)
}
