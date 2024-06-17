package controllers

import (
	"net/http"
	"psy-service/models"
	"psy-service/services"
	"psy-service/utils"

	"github.com/gin-gonic/gin"
)

type PatientController struct {
	Service *services.PatientService
}

func (c *PatientController) CreatePatient(ctx *gin.Context) {
	var patient models.Patient
	if err := ctx.ShouldBindJSON(&patient); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := c.Service.CreatePatient(&patient); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(ctx, http.StatusCreated, patient)
}

func (c *PatientController) UpdatePatient(ctx *gin.Context) {
	id := ctx.Param("id")
	var patient models.Patient
	if err := ctx.ShouldBindJSON(&patient); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := c.Service.UpdatePatient(id, &patient); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(ctx, http.StatusOK, patient)
}

func (c *PatientController) DeletePatient(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.Service.DeletePatient(id); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(ctx, http.StatusOK, gin.H{"result": "success"})
}

func (c *PatientController) GetPatients(ctx *gin.Context) {
	patients, err := c.Service.GetAllPatients()
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(ctx, http.StatusOK, patients)
}
