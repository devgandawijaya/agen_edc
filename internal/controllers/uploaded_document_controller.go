package controllers

import (
	"agen_edc/internal/models"
	"agen_edc/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UploadedDocumentController struct {
	service *services.UploadedDocumentService
}

func NewUploadedDocumentController(s *services.UploadedDocumentService) *UploadedDocumentController {
	return &UploadedDocumentController{service: s}
}

func (c *UploadedDocumentController) Create(ctx *gin.Context) {
	var payload models.UploadedDocument
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

func (c *UploadedDocumentController) Get(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	doc, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	ctx.JSON(http.StatusOK, doc)
}

func (c *UploadedDocumentController) GetByAgentID(ctx *gin.Context) {
	agentIDParam := ctx.Param("agent_id")
	agentID, _ := strconv.Atoi(agentIDParam)
	docs, err := c.service.GetByAgentID(uint(agentID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, docs)
}

func (c *UploadedDocumentController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	var payload models.UploadedDocument
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

func (c *UploadedDocumentController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	if err := c.service.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *UploadedDocumentController) GetAll(ctx *gin.Context) {
	docs, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, docs)
}
