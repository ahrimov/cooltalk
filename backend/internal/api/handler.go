package api

import (
	"database/sql"
	"net/http"

	"github.com/ahrimov/cooltalk-backend/internal/database"
	"github.com/gin-gonic/gin"
)

func getAllUsers(c *gin.Context) {
	users, err := db.GetAllUsers()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func getUserByID(c *gin.Context) {
	userId := c.Param("id")

	user, err := db.GetUserByID(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "user not found", "user_id": userId, "details": err.Error()})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func suggestUsersByUsername(c *gin.Context) {
	username := c.Param("username")

	users, err := db.SuggestUsersByUsername(username)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func addNewUser(c *gin.Context) {
	var user database.User

	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := db.AddNewUser(user)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusAccepted, id)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")

	deletedId, err := db.DeleteUser(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusAccepted, deletedId)
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedData map[string]interface{}

	if err := c.ShouldBindBodyWithJSON(&updatedData); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := db.UpdateUser(id, updatedData)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusAccepted, updatedUser)
}
