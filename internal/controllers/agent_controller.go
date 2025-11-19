package controllers

import (
	"agen_edc/internal/models"
	"agen_edc/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AgentController struct {
	service *services.AgentService
}

func NewAgentController(s *services.AgentService) *AgentController {
	return &AgentController{service: s}
}

func (c *AgentController) Create(ctx *gin.Context) {
	var payload models.Agent
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	actor := ctx.GetHeader("X-Actor") // Assuming actor is passed in header
	if err := c.service.Create(&payload, actor); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, payload)
}

func (c *AgentController) Get(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	agent, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	ctx.JSON(http.StatusOK, agent)
}

func (c *AgentController) GetFull(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	agentFull, err := c.service.GetFullByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	ctx.JSON(http.StatusOK, agentFull)
}

func (c *AgentController) Search(ctx *gin.Context) {
	filters := make(map[string]interface{})
	for key, values := range ctx.Request.URL.Query() {
		if len(values) > 0 {
			filters[key] = values[0]
		}
	}
	agents, err := c.service.Search(filters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, agents)
}

func (c *AgentController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	var payload models.Agent
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payload.ID = uint(id)
	actor := ctx.GetHeader("X-Actor")
	if err := c.service.Update(&payload, actor); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, payload)
}

func (c *AgentController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	actor := ctx.GetHeader("X-Actor")
	if err := c.service.Delete(uint(id), actor); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *AgentController) GetAll(ctx *gin.Context) {
	agents, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, agents)
}

func (c *AgentController) UploadDocuments(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	file, err := ctx.FormFile("document")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// For simplicity, assume we save the file and create a record
	// In real implementation, handle file storage
	doc := &models.UploadedDocument{
		AgentID:  uint(id),
		FileName: file.Filename,
		FilePath: "/uploads/" + file.Filename, // placeholder
	}
	// Assume we have a service for this, but for now, just return success
	ctx.JSON(http.StatusCreated, doc)
}
