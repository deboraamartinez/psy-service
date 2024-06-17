package controllers

import (
	"net/http"
	"psy-service/models"
	"psy-service/services"
	"psy-service/utils"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	Service *services.PaymentService
}

func (c *PaymentController) CreatePayment(ctx *gin.Context) {
	var payment models.Payment
	if err := ctx.ShouldBindJSON(&payment); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := c.Service.CreatePayment(&payment); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(ctx, http.StatusCreated, payment)
}

func (c *PaymentController) UpdatePayment(ctx *gin.Context) {
	id := ctx.Param("id")
	var payment models.Payment
	if err := ctx.ShouldBindJSON(&payment); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := c.Service.UpdatePayment(id, &payment); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(ctx, http.StatusOK, payment)
}

func (c *PaymentController) DeletePayment(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.Service.DeletePayment(id); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(ctx, http.StatusOK, gin.H{"result": "success"})
}

func (c *PaymentController) GetPayments(ctx *gin.Context) {
	payments, err := c.Service.GetAllPayments()
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(ctx, http.StatusOK, payments)
}
