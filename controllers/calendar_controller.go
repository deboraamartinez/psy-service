package controllers

import (
	"net/http"
	"psy-service/models"
	"psy-service/services"
	"psy-service/utils"

	"github.com/gin-gonic/gin"
)

type CalendarController struct {
	Service *services.CalendarService
}

func (c *CalendarController) AddCalendarEvent(ctx *gin.Context) {
	var event models.CalendarEvent
	if err := ctx.ShouldBindJSON(&event); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := c.Service.CreateCalendarEvent(&event); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(ctx, http.StatusCreated, event)
}

func (c *CalendarController) UpdateCalendarEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	var event models.CalendarEvent
	if err := ctx.ShouldBindJSON(&event); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := c.Service.UpdateCalendarEvent(id, &event); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(ctx, http.StatusOK, event)
}

func (c *CalendarController) DeleteCalendarEvent(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.Service.DeleteCalendarEvent(id); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(ctx, http.StatusOK, gin.H{"result": "success"})
}

func (c *CalendarController) GetCalendarEvents(ctx *gin.Context) {
	events, err := c.Service.GetAllCalendarEvents()
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(ctx, http.StatusOK, events)
}
