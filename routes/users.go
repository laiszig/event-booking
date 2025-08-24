package routes

import (
	"event-booking/models"
	"event-booking/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signUpUser(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully."})
}

func userLogin(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse request data."})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate token."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User login successfully.",
		"token":   token,
	})
}
