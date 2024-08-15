package handlers

import (
	"net/http"
	"role-management/internal/api/repository"
	"role-management/internal/models"
	"role-management/pkg/helper"
	"role-management/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

// Create creates a new user
// @Summary Create a new user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User"
// @Router /users [post]
func (h *UserHandler) Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Message(c, http.StatusBadRequest, false, err.Error())
		return
	}

	if err := h.repo.Create(&user); err != nil {
		response.Message(c, http.StatusInternalServerError, false, "Failed to create user")
		return
	}

	response.SendData(c, http.StatusCreated, true, user, nil)
}

// GetByID gets a user by ID
// @Summary Get a user by ID
// @Description Retrieve a user by its ID
// @Tags Users
// @Produce  json
// @Param id path string true "User ID"
// @Router /users/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Message(c, http.StatusBadRequest, false, "Invalid user ID format")
		return
	}

	user, err := h.repo.GetByID(id)
	if err != nil {
		response.Message(c, http.StatusNotFound, false, "User not found")
		return
	}

	response.SendData(c, http.StatusOK, true, user, nil)
}

// Update updates an existing user
// @Summary Update an existing user
// @Description Update a user with the given ID and payload
// @Tags Users
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body models.User true "User"
// @Router /users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Message(c, http.StatusBadRequest, false, "Invalid user ID format")
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Message(c, http.StatusBadRequest, false, err.Error())
		return
	}

	user.ID = id
	if err := h.repo.Update(&user); err != nil {
		response.Message(c, http.StatusInternalServerError, false, "Failed to update user")
		return
	}

	response.SendData(c, http.StatusOK, true, user, nil)
}

// Delete deletes a user by ID
// @Summary Delete a user by ID
// @Description Delete a user by its UUID
// @Tags Users
// @Param id path string true "User ID"
// @Router /users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Message(c, http.StatusBadRequest, false, "Invalid user ID format")
		return
	}

	if err := h.repo.Delete(id); err != nil {
		response.Message(c, http.StatusInternalServerError, false, "Failed to delete user")
		return
	}

	response.Message(c, http.StatusOK, true, "User deleted successfully")
}

// GetAll retrieves all users with pagination
// @Summary Get all users with pagination
// @Description Retrieve a paginated list of users
// @Tags Users
// @Produce  json
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Router /users [get]
func (h *UserHandler) GetAll(c *gin.Context) {
	pagination, _ := helper.GetPagination(c)

	users, totalRecords, err := h.repo.GetAllPaginated(pagination.PageNumber, pagination.PageSize)
	if err != nil {
		response.Message(c, http.StatusInternalServerError, false, "Failed to retrieve users")
		return
	}

	pagination.TotalRecord = totalRecords
	response.SendData(c, http.StatusOK, true, users, pagination)
}
