package app

import (
	"github.com/andrenormanlang/common"
	"github.com/andrenormanlang/database"
	"github.com/andrenormanlang/views"
	"github.com/gin-gonic/gin"
)

// func productHandler(c *gin.Context, app_settings common.AppSettings, db database.Database) ([]byte, error) {
// 	return renderHtml(c, views.MakeProductPage(app_settings.AppNavbar.Links))
// }

func productHandler(c *gin.Context, app_settings common.AppSettings, db database.Database) ([]byte, error) {
	data := map[string]interface{}{
		"title":  "Test Product Generic",
		"slogan": "Save 50%",
	}

	return renderHtml(c, views.MakeProductPage(app_settings.AppNavbar.Links, data))
}