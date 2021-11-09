package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/vshakirova/go-api-project/config"
	"net/http"

	_ "github.com/vshakirova/go-api-project/docs"
)

// @title Project Swagger API
// @version 1.0
// @description Swagger API for Golang Project
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth
func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.Use(config.Auth())
		v1.GET("/users/:id", getUser)
		v1.GET("/users", getUsers)
		v1.PUT("/users/:id", updateUser)
		v1.POST("/users", createUser)
		v1.DELETE("/users/:id", deleteUser)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		return
	}
}
