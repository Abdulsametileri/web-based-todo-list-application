package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

var todoDatabase map[int]Todo

var (
	ErrTodoDescriptionIsEmpty = errors.New("Todo description cannnot be empty.")
)

var (
	MsgTodoAdded = "Todo has been succesfully added."
)

type TodoInput struct {
	TaskDescription string `json:"task_description" binding:"required"`
}

type TodoController interface {
	GetTodoList(ctx *gin.Context)
	AddTodo(ctx *gin.Context)
}

type todoController struct {
	base BaseController
}

func NewTodoController(bctl BaseController) TodoController {
	todoDatabase = make(map[int]Todo, 0)
	return &todoController{
		base: bctl,
	}
}

func (t todoController) GetTodoList(c *gin.Context) {
	panic("implement me")
}

// curl -H "Content-type: application/json" -X POST -d '{"task_description":"dummy"}' localhost:3000/api/v1/addTodo
func (t todoController) AddTodo(c *gin.Context) {
	var vm TodoInput

	if err := c.ShouldBindJSON(&vm); err != nil {
		t.base.Error(c, http.StatusBadRequest, ErrTodoDescriptionIsEmpty)
		return
	}

	todoId := len(todoDatabase)
	addedTodo := Todo{
		Id:          todoId,
		Description: vm.TaskDescription,
	}
	todoDatabase[todoId] = addedTodo

	t.base.Data(c, http.StatusCreated, addedTodo, MsgTodoAdded)
}
