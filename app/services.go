package app

import (
	"github.com/andrenormanlang/go-htmx-mySQL/common"
	"github.com/andrenormanlang/go-htmx-mySQL/database"
	"github.com/andrenormanlang/go-htmx-mySQL/views"
	"github.com/gin-gonic/gin"
)

func servicesHandler(c *gin.Context, app_settings common.AppSettings, db database.Database) ([]byte, error) {
	return renderHtml(c, views.MakeServicesPage(app_settings.AppNavbar.Links))
}
