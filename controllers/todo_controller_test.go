package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTodoController(t *testing.T) {
	bctl := &baseController{}
	todoCtl := NewTodoController(bctl)
	gin.SetMode(gin.TestMode)

	t.Run("[addTodo]: Error when sending empty body", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		todoCtl.AddTodo(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		resBody := Props{}
		json.NewDecoder(w.Body).Decode(&resBody)

		assert.EqualValues(t, ErrTodoDescriptionIsEmpty.Error(), resBody.Message)
	})

	t.Run("[addTodo] Added todo successfully", func(t *testing.T) {
		dummyTaskDescription := "dummy todo"
		reqBody := TodoInput{TaskDescription: dummyTaskDescription}

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		payload, _ := json.Marshal(&reqBody)
		request := httptest.NewRequest(http.MethodPost, "/api/v1/addTodo", bytes.NewBuffer(payload))
		c.Request = request

		todoCtl.AddTodo(c)

		assert.Equal(t, http.StatusCreated, w.Code)

		resBody := Props{}
		json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, MsgTodoAdded, resBody.Message)
		assert.Equal(t, todoDatabase[0].Id, 0)
		assert.Equal(t, todoDatabase[0].Description, dummyTaskDescription)
	})
}
