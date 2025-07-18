package admin_app

import (
	"net/http"

	"github.com/andrenormanlang/go-htmx-mySQL/common"
	"github.com/andrenormanlang/go-htmx-mySQL/database"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(app_settings common.AppSettings, database database.Database) *gin.Engine {

	r := gin.Default()
	r.MaxMultipartMemory = 1

	r.GET("/posts/:id", getPostHandler(database))
	r.POST("/posts", postPostHandler(database))
	r.PUT("/posts", putPostHandler(database))
	r.DELETE("/posts/:id", deletePostHandler(database))

	r.POST("/images", postImageHandler(app_settings))
	r.DELETE("/images/:name", deleteImageHandler(app_settings))

	r.POST("/pages", postPageHandler(database))

	// ASSUME: permalink will not contain /
	// POST gocms.com/permalinks/my/permalink
	r.POST("/permalinks/:permalink/:post_id", postPermalinkHandler(database))

	// For container health purposes
	r.Any("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, PostIdResponse{Id: 0})
	})

	r.GET("/cards/:schema", getCardHandler(database))
	r.GET("/cards/:schema/:limit/:page", getCardHandler(database))
	r.POST("/cards", postCardHandler(database))
	r.POST("/card-schemas", postSchemaHandler(database))

	return r
}
