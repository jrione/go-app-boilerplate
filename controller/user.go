package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jrione/go-app-boilerplate/helper"
	"github.com/jrione/go-app-boilerplate/plugin"
	"github.com/jrione/go-app-boilerplate/repository"
)

type UserController struct {
	userRepo repository.UserRepository
	logger   *plugin.Logger
}

func NewUserController(userRepo repository.UserRepository, logger *plugin.Logger) *UserController {
	return &UserController{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (uc *UserController) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		uc.logger.Error("Invalid user ID: ", err)
		helper.ErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := uc.userRepo.GetUserByID(uint(id))
	if err != nil {
		uc.logger.Error("User not found: ", err)
		helper.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	helper.JSONResponse(c, http.StatusOK, user)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user repository.User
	if err := c.ShouldBindJSON(&user); err != nil {
		uc.logger.Error("Invalid JSON: ", err)
		helper.ErrorResponse(c, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if err := uc.userRepo.CreateUser(&user); err != nil {
		uc.logger.Error("Failed to create user: ", err)
		helper.ErrorResponse(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	helper.JSONResponse(c, http.StatusCreated, user)
}
