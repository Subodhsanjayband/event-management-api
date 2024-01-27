package routes

import (
	"net/http"

	"github.com/Subodhsanjayband/event_manager/models"
	"github.com/Subodhsanjayband/event_manager/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the request data.",
		})
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create the user.",
		})
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "Created new user.", "user": user,
	})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse the request data.",
		})
	}
	err = user.ValidateUser()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}
	id, err := user.GetUserID()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	user.ID = id
	token, err := utils.GenerateTokan(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful!!!", "token": token, "user": user,
	})
}

func getUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could fetch users.",
		})
		return
	}
	c.JSON(http.StatusOK, users)

}
