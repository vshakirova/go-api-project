package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vshakirova/go-api-project/models"
	"net/http"
	"strconv"
)

var users = map[string]models.User{
	"1": {Name: "John Doe", Address: "N street", Job: "worker"},
	"2": {Name: "Steff Mur", Address: "J street", Job: "stuff"},
}

// getUsers godoc
// @Summary Retrieves all users
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
// @Security BasicAuth
func getUsers(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, users)
}

// getUser godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {array} models.User
// @Router /users/{id} [get]
// @Security BasicAuth
func getUser(ctx *gin.Context) {
	id := ctx.Param("id")

	for key, user := range users {
		if key == id {
			ctx.IndentedJSON(http.StatusOK, user)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "user not found"})
}

// updateUser godoc
// @Summary Update user info based on given ID
// @Produce json
// @Accept json
// @Param id path integer true "User ID"
// @Param data body models.User true "User Info"
// @Success 200 {object} models.User
// @Router /users/{id} [put]
// @Security BasicAuth
func updateUser(ctx *gin.Context) {
	var newUser models.User
	id := ctx.Param("id")

	if err := ctx.BindJSON(&newUser); err != nil {
		return
	}
	if _, ok := users[id]; ok != false {
		users[id] = newUser
		return
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "user not found"})
}

// deleteUser godoc
// @Summary Delete user by ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200
// @Router /users/{id} [delete]
// @Security BasicAuth
func deleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	if _, ok := users[id]; ok != false {
		delete(users, id)
		return
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "user not found"})
}

// createUser godoc
// @Summary Creates user based on user info
// @Produce json
// @Accept json
// @Param data body models.User true "User Info"
// @Success 201 {object} models.User
// @Router /users [post]
// @Security BasicAuth
func createUser(ctx *gin.Context) {
	var newUser models.User

	if err := ctx.BindJSON(&newUser); err != nil {
		return
	}

	users[strconv.FormatInt(int64(len(users)+1), 10)] = newUser
	ctx.IndentedJSON(http.StatusCreated, newUser)
}
