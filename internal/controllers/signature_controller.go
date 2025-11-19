package controllers

import (
	"agen_edc/internal/models"
	"agen_edc/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SignatureController struct {
	service *services.SignatureService
}

func NewSignatureController(s *services.SignatureService) *SignatureController {
	return &SignatureController{service: s}
}

func (c *SignatureController) Create(ctx *gin.Context) {
	var payload models.Signature
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

func (c *SignatureController) Get(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	sig, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	ctx.JSON(http.StatusOK, sig)
}

func (c *SignatureController) GetByAgentID(ctx *gin.Context) {
	agentIDParam := ctx.Param("agent_id")
	agentID, _ := strconv.Atoi(agentIDParam)
	sigs, err := c.service.GetByAgentID(uint(agentID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, sigs)
}

func (c *SignatureController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	var payload models.Signature
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

func (c *SignatureController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	if err := c.service.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *SignatureController) GetAll(ctx *gin.Context) {
	sigs, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, sigs)
}
