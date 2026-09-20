package handler

import (
	"backend/auth"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userServices *services.UserService
}

func NewUserHanlder(userServices *services.UserService) *UserHandler {
	return &UserHandler{
		userServices: userServices,
	}
}
func (h *UserHandler) Register(c *gin.Context) {

	var input services.RegisterUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	user, err := h.userServices.Register(
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
		"message": "user registered successfully",
		"user":    user,
	})
}

// POST /auth/login
func (h *UserHandler) Login(c *gin.Context) {

	var input services.LoginUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	user, err := h.userServices.Login(
		c.Request.Context(),
		input,
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"user":    user,
		"token":   token,
	})
}
