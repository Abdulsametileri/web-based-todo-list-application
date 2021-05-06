package controllers

import "github.com/gin-gonic/gin"

type TodoController interface {
	GetTodoList(ctx *gin.Context)
	AddTodo(ctx *gin.Context)
}

type todoController struct{}

func NewTodoController() TodoController {
	return &todoController{}
}

func (t todoController) GetTodoList(c *gin.Context) {
	panic("implement me")
}

func (t todoController) AddTodo(c *gin.Context) {
	panic("implement me")
}
