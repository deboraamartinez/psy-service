package controllers

import (
	"net/http"
	"psy-service/models"
	"psy-service/services"
	"psy-service/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service *services.AuthService
}

func (c *AuthController) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := c.Service.Register(&user); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(ctx, http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	token, err := c.Service.Login(&user)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	utils.RespondWithJSON(ctx, http.StatusOK, gin.H{"token": token})
}
