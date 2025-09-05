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
	router.GET("/users", getAllUsers)
}
