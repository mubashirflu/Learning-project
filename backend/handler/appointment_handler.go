package handler

import (
	"net/http"
	"strconv"
	"time"

	"backend/services"

	"github.com/gin-gonic/gin"
)

type AppointmentHandler struct {
	appointmentService *services.AppointmentService
}

func NewAppointmentHandler(
	appointmentService *services.AppointmentService,
) *AppointmentHandler {
	return &AppointmentHandler{
		appointmentService: appointmentService,
	}
}

// POST /api/appointments
func (h *AppointmentHandler) Create(c *gin.Context) {

	var input services.CreateAppointmentInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	appointment, err := h.appointmentService.CreateAppointment(
		c.Request.Context(),
		input,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "appointment created successfully",
		"appointment": appointment,
	})
}

// GET /api/appointments/customer/:customer_id
func (h *AppointmentHandler) GetByCustomer(c *gin.Context) {

	customerID, err := strconv.ParseUint(
		c.Param("customer_id"),
		10,
		32,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid customer id",
		})
		return
	}

	appointments, err := h.appointmentService.GetAppointmentsByCustomer(
		c.Request.Context(),
		uint(customerID),
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"appointments": appointments,
	})
}

// GET /api/appointments/service/:service_id
func (h *AppointmentHandler) GetByService(c *gin.Context) {

	serviceID, err := strconv.ParseUint(
		c.Param("service_id"),
		10,
		32,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid service id",
		})
		return
	}

	appointments, err := h.appointmentService.GetAppointmentsByService(
		c.Request.Context(),
		uint(serviceID),
		time.Now(),
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"appointments": appointments,
	})
}

// PUT /api/appointments/:id/cancel
func (h *AppointmentHandler) Cancel(c *gin.Context) {

	appointmentID, err := strconv.ParseUint(
		c.Param("id"),
		10,
		32,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid appointment id",
		})
		return
	}

	err = h.appointmentService.CancelAppointment(
		c.Request.Context(),
		uint(appointmentID),
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "appointment cancelled successfully",
	})
}
func (h *AppointmentHandler) GetByDate(c *gin.Context) {

	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not found",
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

	dateParam := c.Param("date")

	appointmentDate, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid date format, use YYYY-MM-DD",
		})
		return
	}

	appointments, err := h.appointmentService.GetAppointmentsByDate(
		c.Request.Context(),
		userID,
		appointmentDate,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"appointments": appointments,
	})
}
