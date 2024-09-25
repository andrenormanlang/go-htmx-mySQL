package admin_app

import (
	"github.com/gin-gonic/gin"
	"github.com/andrenormanlang/common"
	"github.com/andrenormanlang/database"
)

func SetupRoutes(app_settings common.AppSettings, database database.Database) *gin.Engine {

	r := gin.Default()
	// r.Run(":8081")
	r.MaxMultipartMemory = 1

	r.GET("/posts/:id", getPostHandler(database))
	r.POST("/posts", postPostHandler(database))
	r.PUT("/posts", putPostHandler(database))
	r.DELETE("/posts/:id", deletePostHandler(database))

	r.POST("/images", postImageHandler(app_settings))
	r.DELETE("/images/:name", deleteImageHandler(app_settings))

	r.POST("/pages", postPageHandler(database))

	return r
}