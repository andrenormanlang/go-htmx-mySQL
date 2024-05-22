package app

import (
	"net/http"
	"github.com/a-h/templ"
	"github.com/andrenormanlang/database"
	views "github.com/andrenormanlang/views/index"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type AppSettings struct {
    Database_address string
    Database_port  int
	Database_user string
	Database_password string
}

func Run(app_settings AppSettings, database database.Database) error {
	r := gin.Default()
	r.MaxMultipartMemory = 1
	//r.LoadHTMLFiles("./templates/contact/contact-success.html", "./templates/contact/contact-failure.html")

	r.GET("/", makeHomeHandler(app_settings, database))

	// Contact form related endpoints
	r.GET("/contact", makeContactPageHandler(app_settings, database))
	r.POST("/contact-send", makeContactFormHandler())

	// Post related endpoints
	r.GET("/post/:id", makePostHandler(database))


	r.Static("/static", "./static")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	return nil
}

func Render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

/// This function will act as the handler for
/// the home page
func makeHomeHandler(settings AppSettings, db database.Database) func(*gin.Context) {
	return func(c *gin.Context){
		posts, err := db.GetPosts()
		if err != nil {
			log.Error().Msgf("error loading posts: %v\n", err)
			return
		}

		Render(c, http.StatusOK, views.MakeIndex(posts))
		
	}
}