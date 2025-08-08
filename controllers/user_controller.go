package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/credondocr/github-workflow-showcase/models"
)

// UserController handles HTTP requests related to users.
type UserController struct {
	userRepo models.UserRepository
}

// NewUserController creates a new instance of the user controller.
func NewUserController(userRepo models.UserRepository) *UserController {
	return &UserController{
		userRepo: userRepo,
	}
}

// GetUsers handles GET /users - retrieves all users.
func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.userRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal server error",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    users,
		"total":   len(users),
	})
}

// GetUser handles GET /users/:id - retrieves a user by ID.
func (uc *UserController) GetUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid ID",
			"message": "ID must be an integer",
		})

		return
	}

	user, err := uc.userRepo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "User not found",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})
}

// CreateUser handles POST /users - creates a new user.
func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid data",
			"message": err.Error(),
		})

		return
	}

	if err := user.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"message": err.Error(),
		})

		return
	}

	if err := uc.userRepo.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error creating user",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User created successfully",
		"data":    user,
	})
}

// UpdateUser handles PUT /users/:id - updates an existing user.
func (uc *UserController) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid ID",
			"message": "ID must be an integer",
		})

		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid data",
			"message": err.Error(),
		})

		return
	}

	if err := user.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"message": err.Error(),
		})

		return
	}

	if err := uc.userRepo.Update(id, &user); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "User not found",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User updated successfully",
		"data":    user,
	})
}

// DeleteUser handles DELETE /users/:id - deletes a user.
func (uc *UserController) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid ID",
			"message": "ID must be an integer",
		})

		return
	}

	if err := uc.userRepo.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "User not found",
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User deleted successfully",
	})
}

// HealthCheck handles GET /health - health check endpoint.
func (uc *UserController) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "API is working correctly",
		"version": "1.0.0",
	})
}

// GetStats handles GET /api/v1/users/stats - get user statistics.
func (uc *UserController) GetStats(c *gin.Context) {
	users, err := uc.userRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to retrieve users",
			"message": err.Error(),
		})

		return
	}

	totalUsers := len(users)

	// Calculate age statistics
	var totalAge int

	var minAge, maxAge int

	ageRanges := map[string]int{
		"18-25": 0,
		"26-35": 0,
		"36-50": 0,
		"51+":   0,
	}

	if totalUsers > 0 {
		minAge = users[0].Age
		maxAge = users[0].Age

		for _, user := range users {
			totalAge += user.Age

			if user.Age < minAge {
				minAge = user.Age
			}

			if user.Age > maxAge {
				maxAge = user.Age
			}

			// Categorize by age range
			switch {
			case user.Age >= 18 && user.Age <= 25:
				ageRanges["18-25"]++
			case user.Age >= 26 && user.Age <= 35:
				ageRanges["26-35"]++
			case user.Age >= 36 && user.Age <= 50:
				ageRanges["36-50"]++
			default:
				ageRanges["51+"]++
			}
		}
	}

	var averageAge float64
	if totalUsers > 0 {
		averageAge = float64(totalAge) / float64(totalUsers)
	}

	stats := gin.H{
		"total_users":  totalUsers,
		"average_age":  averageAge,
		"min_age":      minAge,
		"max_age":      maxAge,
		"age_ranges":   ageRanges,
		"generated_at": time.Now().Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User statistics retrieved successfully",
		"data":    stats,
	})
}
