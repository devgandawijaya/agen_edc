package controllers

import (
	"agen_edc/internal/models"
	"agen_edc/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AcquisitionController struct {
	service *services.AcquisitionService
}

func NewAcquisitionController(s *services.AcquisitionService) *AcquisitionController {
	return &AcquisitionController{service: s}
}

func (c *AcquisitionController) Create(ctx *gin.Context) {
	var payload models.AcquisitionInfo
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.Create(&payload); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, payload)
}

func (c *AcquisitionController) Get(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	acquisition, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	ctx.JSON(http.StatusOK, acquisition)
}

func (c *AcquisitionController) GetByAgentID(ctx *gin.Context) {
	agentIDParam := ctx.Param("agent_id")
	agentID, _ := strconv.Atoi(agentIDParam)
	acquisition, err := c.service.GetByAgentID(uint(agentID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	ctx.JSON(http.StatusOK, acquisition)
}

func (c *AcquisitionController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	var payload models.AcquisitionInfo
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payload.ID = uint(id)
	if err := c.service.Update(&payload); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, payload)
}

func (c *AcquisitionController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	if err := c.service.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *AcquisitionController) GetAll(ctx *gin.Context) {
	acquisitions, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, acquisitions)
}
