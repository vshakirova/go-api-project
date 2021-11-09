package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type user struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Job     string `json:"job"`
}

var users = map[string]user{
	"1": { Name: "John Doe", Address: "N street", Job: "worker"},
	"2": {Name: "Steff Mur", Address: "J street", Job: "stuff"},
}

// getUsers godoc
// @Summary Retrieves all users
// @Produce json
// @Success 200 {array} user
// @Router /users [get]
// @Security BasicAuth
func getUsers(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, users)
}

// getUser godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {array} user
// @Router /users/{id} [get]
// @Security BasicAuth
func getUser(ctx *gin.Context ) {
	id := ctx.Param("id")

	for key, user := range users {
		if key == id {
			ctx.IndentedJSON(http.StatusOK, user)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

// updateUser godoc
// @Summary Update user info based on given ID
// @Produce json
// @Accept json
// @Param id path integer true "User ID"
// @Param data body user true "User Info"
// @Success 200 {object} user
// @Router /users/{id} [put]
// @Security BasicAuth
func updateUser(ctx *gin.Context) {
	var newUser user
	id := ctx.Param("id")

	if err := ctx.BindJSON(&newUser); err != nil {
		return
	}
	if _, ok := users[id]; ok!=false {
		users[id] = newUser
		return
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

// deleteUser godoc
// @Summary Delete user by ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200
// @Router /users [delete]
// @Security BasicAuth
func deleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	if _, ok := users[id]; ok!=false {
		delete(users, id)
		return
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

// createUser godoc
// @Summary Creates user based on user info
// @Produce json
// @Accept json
// @Param data body user true "User Info"
// @Success 200 {object} user
// @Router /users [post]
// @Security BasicAuth
func createUser(ctx *gin.Context) {
	var newUser user

	if err := ctx.BindJSON(&newUser); err != nil {
		return
	}

	users[strconv.FormatInt(int64(len(users)+1), 10)] = newUser
	ctx.IndentedJSON(http.StatusCreated, newUser)
}
