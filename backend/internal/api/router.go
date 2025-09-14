package api

import (
	"github.com/ahrimov/cooltalk-backend/internal/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var db *database.MainDB

func SetUpRouter(mainDB *database.MainDB) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	db = mainDB

	setEndpoints(router)

	return router
}

func setEndpoints(router *gin.Engine) {
	mainAPI := router.Group("/api")

	v1 := mainAPI.Group("/v1")
	v1.GET("/users", getAllUsers)
	v1.GET("/user/:id", getUserByID)
	v1.GET("/users/suggest/:username", suggestUsersByUsername)

	v1.POST("/user", addNewUser)

	v1.PUT("/user/:id", updateUser)

	v1.DELETE("/user/:id", deleteUser)
}
