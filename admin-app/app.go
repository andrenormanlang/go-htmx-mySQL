package admin_app

import (
	// "log"

	"github.com/andrenormanlang/common"
	"github.com/andrenormanlang/database"
	"github.com/gin-gonic/gin"
	// lua "github.com/yuin/gopher-lua"
)


// func SetupRoutes(app_settings common.AppSettings, database database.Database, plugins map[string]*lua.LState) *gin.Engine {

	// r := gin.Default()
	// // r.Run(":8081")
	// r.MaxMultipartMemory = 1

// 	post_hook, ok := hooks["add_post"]
// 	if !ok {
// 		log.Fatalf("Could not find hook for add_post")
// 	}


// 	r.GET("/posts/:id", getPostHandler(database))
// 	r.POST("/posts", postPostHandler(database, shortcode_handlers,post_hook.(plugins.PostHook)))
// 	r.PUT("/posts", putPostHandler(database))
// 	r.DELETE("/posts/:id", deletePostHandler(database))

// 	r.POST("/images", postImageHandler(app_settings))
// 	r.DELETE("/images/:name", deleteImageHandler(app_settings))

// 	r.POST("/pages", postPageHandler(database))

// 	return r
// }

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

	r.POST("/card-schemas", postSchemaHandler(database))
	r.POST("/cards", postCardHandler(database))

	return r
}