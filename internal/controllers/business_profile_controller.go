package controllers

import (
	"agen_edc/internal/models"
	"agen_edc/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BusinessProfileController struct {
	service *services.BusinessProfileService
}

func NewBusinessProfileController(s *services.BusinessProfileService) *BusinessProfileController {
	return &BusinessProfileController{service: s}
}

func (c *BusinessProfileController) Create(ctx *gin.Context) {
	var payload models.BusinessProfile
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

func (c *BusinessProfileController) Get(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	profile, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	ctx.JSON(http.StatusOK, profile)
}

func (c *BusinessProfileController) GetByAgentID(ctx *gin.Context) {
	agentIDParam := ctx.Param("agent_id")
	agentID, _ := strconv.Atoi(agentIDParam)
	profiles, err := c.service.GetByAgentID(uint(agentID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, profiles)
}

func (c *BusinessProfileController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	var payload models.BusinessProfile
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

func (c *BusinessProfileController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	if err := c.service.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *BusinessProfileController) GetAll(ctx *gin.Context) {
	profiles, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, profiles)
}
