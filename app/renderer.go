package app

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

// This function will render the templ component into 
//  a gin context´s Response Writer
func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}