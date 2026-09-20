package handler

import (
	"net/http"
	"strconv"

	"backend/services"

	"github.com/gin-gonic/gin"
)

type ServiceHandler struct {
	serviceService *services.ServiceService
}

func NewServiceHandler(
	serviceService *services.ServiceService,
) *ServiceHandler {
	return &ServiceHandler{
		serviceService: serviceService,
	}
}

// POST /api/services
func (h *ServiceHandler) Create(c *gin.Context) {

	// Get user_id from JWT middleware
	userIDValue, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	userID, ok := userIDValue.(uint)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user id",
		})
		return
	}

	// Request body
	var input services.CreateServiceInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	// Call service layer
	newService, err := h.serviceService.CreateService(
		c.Request.Context(),
		userID,
		input,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "service created successfully",
		"service": newService,
	})
}

// GET /api/services
func (h *ServiceHandler) GetAll(c *gin.Context) {

	userIDValue, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	userID, ok := userIDValue.(uint)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user id",
		})
		return
	}

	// Get user's services
	serviceList, err := h.serviceService.GetServices(
		c.Request.Context(),
		userID,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"services": serviceList,
	})
}

// GET /api/services/:id
func (h *ServiceHandler) GetByID(c *gin.Context) {

	userIDValue, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	userID, ok := userIDValue.(uint)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user id",
		})
		return
	}

	// Get service ID from URL
	serviceID, err := strconv.ParseUint(
		c.Param("id"),
		10,
		32,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid service id",
		})
		return
	}

	// Get service
	service, err := h.serviceService.GetService(
		c.Request.Context(),
		userID,
		uint(serviceID),
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"service": service,
	})
}

// PUT /api/services/:id
func (h *ServiceHandler) Update(c *gin.Context) {

	userIDValue, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	userID, ok := userIDValue.(uint)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user id",
		})
		return
	}

	// Get service ID
	serviceID, err := strconv.ParseUint(
		c.Param("id"),
		10,
		32,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid service id",
		})
		return
	}

	// Request body
	var input services.UpdateServiceInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	// Update service
	updatedService, err := h.serviceService.UpdateService(
		c.Request.Context(),
		userID,
		uint(serviceID),
		input,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "service updated successfully",
		"service": updatedService,
	})
}

// DELETE /api/services/:id
func (h *ServiceHandler) Delete(c *gin.Context) {

	userIDValue, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	userID, ok := userIDValue.(uint)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user id",
		})
		return
	}

	// Get service ID
	serviceID, err := strconv.ParseUint(
		c.Param("id"),
		10,
		32,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid service id",
		})
		return
	}

	// Delete service
	err = h.serviceService.DeleteService(
		c.Request.Context(),
		userID,
		uint(serviceID),
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "service deleted successfully",
	})
}
