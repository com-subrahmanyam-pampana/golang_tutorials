package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID   int    `json:"id"`
	NAME string `json:"name"`
}

var users = []user{
	{ID: 1, NAME: "subbu"}, {ID: 2, NAME: "trump"},
}

func getUsers(ginContext *gin.Context) {
	ginContext.JSON(http.StatusOK, users)
}

func getWithStruct() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.Run("localhost:9091")
}
