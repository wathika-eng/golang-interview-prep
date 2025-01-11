package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"github.com/golodash/galidator"
	"github.com/matthewjamesboyle/golang-interview-prep/internal/db"
	"github.com/matthewjamesboyle/golang-interview-prep/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var (
	g           = galidator.New()
	customizer  = g.Validator(models.User{})
	ctx         = context.Background()
	redisClient = db.RedisInit()
)

// test if API is up
func Test(c *gin.Context) {
	c.JSON(200, gin.H{"message": "server is up and running"})
}

// CreateUser handles user creation and stores the user in the database
func CreateUser(c *gin.Context) {
	var user models.User
	// Bind incoming JSON to the user struct
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"error":       customizer.DecryptErrors(err),
		})
		return
	}

	// Hash the password before saving to the database
	hashedPassword := hashpass(user.Password)
	user.Password = hashedPassword
	user.ID = uuid.Must(uuid.NewV7())
	// user.WorkID, _ = strconv.Atoi(string(user.WorkID))
	fmt.Println(user.WorkID)
	userID, err := models.CreateUser(&user)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status_code": http.StatusBadRequest,
			"error":       err.Error(),
		})
		return
	}
	// Return success response with user ID and email
	c.JSON(http.StatusCreated, gin.H{
		"message":    "User created successfully",
		"user_email": user.Email,
		"id":         userID,
	})
}

// fetch a user by ID
func GetUserByID(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status_code": http.StatusBadRequest,
			"error":       err.Error(),
		})
		return
	}
	user, err := models.GetUser(ID)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status_code": http.StatusBadRequest,
			"error":       err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status_code": 200,
		"message":     user,
	})
}

// fetch all users from the DB
func GetUsers(c *gin.Context) {
	val, err := redisClient.Get(ctx, "users").Result()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status_code": http.StatusOK,
			"users":       val,
		})
		return
	}
	users, err := models.GetUsers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"error":       "Failed to fetch users from database: " + err.Error(),
		})
		return
	}
	userData, err := json.Marshal(users)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"error":       "Failed to marshal users: " + err.Error(),
		})
		return
	}
	err = redisClient.Set(ctx, "users", userData, 100000).Err()
	if err != nil {
		log.Printf("failed to cache users: %v\n", err)
	}

	// Return the users
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     users,
	})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	workID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "Invalid user ID format",
		})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     customizer.DecryptErrors(err),
		})
		return
	}

	updatedUser, err := models.UpdateUser(workID, user.UserName, user.Email, user.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": http.StatusNotFound,
			"message":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     fmt.Sprintf("User with ID %d updated successfully", workID),
		"user":        updatedUser,
	})
}

func DeleteUser(c *gin.Context) {
	workID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "invalid ID format",
		})
		return
	}

	err = models.DeleteUser(workID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"message":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     fmt.Sprintf("user with ID %d deleted successfully", workID),
	})
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status_code": http.StatusNotFound,
		"message":     "route not found",
	})
	// c.Redirect(200, routes.)
}

// hashpass hashes the user's password using bcrypt
func hashpass(pass string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}
	return string(hashedPassword)
}
