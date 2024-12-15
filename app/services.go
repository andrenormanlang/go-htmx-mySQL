package app
import (
	"github.com/gin-gonic/gin"
	"github.com/andrenormanlang/common"
	"github.com/andrenormanlang/database"
	"github.com/andrenormanlang/views"
)
func servicesHandler(c *gin.Context, app_settings common.AppSettings, db database.Database) ([]byte, error) {
	return renderHtml(c, views.MakeServicesPage(app_settings.AppNavbar.Links))
}