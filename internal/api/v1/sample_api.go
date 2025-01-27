package v1

import (
	"go-api-project-seed/internal/model"
	"go-api-project-seed/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SampleAPI handles API operations for Sample.
type SampleAPI struct {
	service *service.SampleService
}

// NewSampleAPI initializes a new instance of SampleAPI.
func NewSampleAPI(service *service.SampleService) *SampleAPI {
	return &SampleAPI{service: service}
}

// RegisterRoutes registers API routes for Sample.
func (api *SampleAPI) RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("sample")
	{
		group.GET("/", api.GetAll)
		group.GET("/:id", api.GetByID)
		group.POST("/", api.Create)
		group.PUT("/:id", api.Update)
		group.DELETE("/:id", api.Delete)
	}
}

// GetAll retrieves all Sample entries.
func (api *SampleAPI) GetAll(c *gin.Context) {
	data, err := api.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

// GetByID retrieves a single Sample entry by ID.
func (api *SampleAPI) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	data, err := api.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

// Create adds a new Sample entry.
func (api *SampleAPI) Create(c *gin.Context) {
	var payload model.Sample
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := api.service.Create(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": payload})
}

// Update modifies an existing Sample entry.
func (api *SampleAPI) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var payload model.Sample
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payload.Id = strconv.Itoa(int(uint(id)))
	if err := api.service.Update(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": payload})
}

// Delete removes an existing Sample entry.
func (api *SampleAPI) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := api.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
