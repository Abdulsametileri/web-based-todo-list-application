package main

import (
	"context"
	"github.com/Abdulsametileri/web-based-todo-list-application/config"
	"github.com/Abdulsametileri/web-based-todo-list-application/controllers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	config.Setup()

	if !config.IsDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	var router = gin.New()

	if config.IsDebug {
		router.Use(gin.Logger())
	}

	baseCtl := controllers.NewBaseController()
	todoCtl := controllers.NewTodoController(baseCtl)

	router.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

	v1 := router.Group("api/v1")
	{
		v1.GET("getTodoList", todoCtl.GetTodoList)
		v1.POST("addTodo", todoCtl.AddTodo)
	}

	srv := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got signal:", sig)
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	srv.Shutdown(ctx)
}

func todos(c *gin.Context) {

}

func todo(c *gin.Context) {

}
