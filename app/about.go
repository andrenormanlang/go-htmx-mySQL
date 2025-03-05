package app

import (
	"github.com/andrenormanlang/common"
	"github.com/andrenormanlang/database"
	"github.com/andrenormanlang/views"
	"github.com/gin-gonic/gin"
)

func aboutHandler(c *gin.Context, app_settings common.AppSettings, db database.Database) ([]byte, error) {
	return renderHtml(c, views.MakeAboutPage(app_settings.AppNavbar.Links))
}
