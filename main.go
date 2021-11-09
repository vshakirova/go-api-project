package main

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"strings"

	_ "github.com/vshakirova/go-api-project/docs"
)

// @title Project Swagger API
// @version 1.0
// @description Swagger API for Golang Project

// @BasePath /api/v1
// @securityDefinitions.basic
// @in header
// @name Authorization

// @Security BasicAuth

var loginPassword = "mirantis:mirantis"

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.Use(auth())
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

func auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		fmt.Println(authHeader)

		if len(authHeader) == 0 {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "header is missed"})
			ctx.Abort()
		}
		decodedBasicAuth := parseBasicAuth(authHeader)

		if decodedBasicAuth != loginPassword {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "this user isn't authorized to this operation"})
			ctx.Abort()
		}
		ctx.Next()
	}
}

func parseBasicAuth(auth string) string {
	str := strings.SplitN(auth, " ", 2)
	if len(str) != 2 {
		fmt.Println("failed to parse authentication string")
	}
	if str[0] != "Basic" {
		fmt.Println("authorization scheme is not Basic")
	}
	c, err := base64.StdEncoding.DecodeString(str[1])
	if err != nil {
		fmt.Println("failed to parse base64 credentials")
	}
	return string(c)
}
