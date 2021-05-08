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
	if !config.IsDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	var router = gin.New()

	if config.IsDebug {
		router.Use(gin.Logger())
	}

	baseCtl := controllers.NewBaseController()
	todoCtl := controllers.NewTodoController(baseCtl)

	router.Use(corsMiddleware())

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

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.IsDebug {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "http://ec2-52-59-226-235.eu-central-1.compute.amazonaws.com:8080")
			c.Writer.Header().Set("Access-Control-Allow-Origin", "http://ec2-3-126-245-149.eu-central-1.compute.amazonaws.com:8080")
		}

		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
