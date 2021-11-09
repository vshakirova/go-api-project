package config

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var loginPassword = "mirantis:mirantis"

func Auth() gin.HandlerFunc {
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
