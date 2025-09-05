package api

import (
	"net/http"

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
