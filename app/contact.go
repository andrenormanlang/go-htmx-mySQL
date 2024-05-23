package app

import (
	"net/http"
	"net/mail"

	"github.com/andrenormanlang/database"
	"github.com/andrenormanlang/views"
	"github.com/gin-gonic/gin"
)

func makeContactFormHandler() func(*gin.Context) {
	return func(c *gin.Context) {
		c.Request.ParseForm()
		email := c.Request.FormValue("email")
		name := c.Request.FormValue("name")
		message := c.Request.FormValue("message")

		// Parse email
		_, err := mail.ParseAddress(email)
		if err != nil { 
			render(c, http.StatusBadRequest, views.MakeContactFailure(email, err.Error()))
			return	
		}

		// Make sure name and message is reasonable
		if len(name) > 200 {
			render(c, http.StatusBadRequest, views.MakeContactFailure(email, err.Error()))
			return
		} 

		if len(message) > 10000 {
			render(c, http.StatusBadRequest, views.MakeContactFailure(email, err.Error()))
			return
		}

		render(c, http.StatusOK, views.MakeContactSuccess(email, name))
		c.HTML(http.StatusOK, "contact-success.html", gin.H{
			"name": name,
			"email": email,
		})
	}
}

// TODO : This is a duplicate of the index handler... abstract
func makeContactPageHandler(settings AppSettings, db database.Database) func(*gin.Context) {
	return func(c *gin.Context){
		render(c, http.StatusOK, views.MakeContactPage())
	}
}