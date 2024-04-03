package handler

import (
	"github.com/gin-gonic/gin"
	"mygram-railway/service"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) FindOneUser(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Malformat ID",
		})
		return
	}

	res, err := h.userService.FindOneUser(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Failed to find product",
			"detail":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
