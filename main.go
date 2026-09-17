package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = connectDatabase()

func main() {

	router := gin.Default()

	// Create/update database table based on User model
	db.AutoMigrate(&User{})

	// Home
	router.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to my Gin API",
		})

	})

	// CREATE
	router.POST("/users", createUser)

	// READ ALL
	router.GET("/users", getUsers)

	// READ ONE
	router.GET("/users/:id", getUser)

	// UPDATE
	router.PUT("/users/:id", updateUser)

	// DELETE
	router.DELETE("/users/:id", deleteUser)

	router.Run(":4000")
}

// CREATE USER
func createUser(c *gin.Context) {

	var user User

	err := c.ShouldBindJSON(&user)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
		})

		return
	}

	result := db.Create(&user)

	if result.Error != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	c.JSON(http.StatusCreated, user)
}

// GET ALL USERS
func getUsers(c *gin.Context) {

	var users []User

	result := db.Find(&users)

	if result.Error != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch users",
		})

		return
	}

	c.JSON(http.StatusOK, users)
}

// GET SINGLE USER
func getUser(c *gin.Context) {

	id := c.Param("id")

	var user User

	result := db.First(&user, "id = ?", id)

	if result.Error != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})

		return
	}

	c.JSON(http.StatusOK, user)
}

// UPDATE USER
func updateUser(c *gin.Context) {

	id := c.Param("id")

	var user User

	// Find existing user
	result := db.First(&user, "id = ?", id)

	if result.Error != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})

		return
	}

	// Receive new data
	var input User

	err := c.ShouldBindJSON(&input)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
		})

		return
	}

	// Update user information
	user.Name = input.Name

	// Save changes
	result = db.Save(&user)

	if result.Error != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user",
		})

		return
	}

	c.JSON(http.StatusOK, user)
}

// DELETE USER
func deleteUser(c *gin.Context) {

	id := c.Param("id")

	var user User

	// Find existing user
	result := db.First(&user, "id = ?", id)

	if result.Error != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})

		return
	}

	// Delete user from database
	result = db.Delete(&user)

	if result.Error != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
