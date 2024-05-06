package main

import (
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc *UserController) GetUserInfo(c *gin.Context) {
	userID := c.Param("id")
	c.JSON(200, gin.H{"id": userID, "name": "panda", "email": "liamhu0914@gmail.com"})
}

func main() {
	router := gin.Default()

	UserController := &UserController{}
	router.GET("/user/:id", UserController.GetUserInfo)

	router.Run(":8080")
}
