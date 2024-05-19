package main

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.MaxMultipartMemory = 1
  r.LoadHTMLFiles(
    "./templates/contact/contact-success.html",
    "./templates/contact/contact-failure.html",
  )
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

  // Contact form endpoint
  r.POST("/contact-send", func(c *gin.Context) {
    c.Request.ParseForm()
    name := c.Request.FormValue("name")
    email := c.Request.FormValue("email")
    message := c.Request.FormValue("message")

    // Parse email and make sure it is valid
    _, err := mail.ParseAddress(email)
    if err != nil {
      c.HTML(http.StatusOK, "contact-failure.html", gin.H{
        "email": email,
        "error": "Invalid email",
      })
      return
    }


    // Name and message validation
    if len(name) > 50 {
      c.HTML(http.StatusOK, "contact-failure.html", gin.H{
        "email": email,
        "error": "Invalid name",
      })
      return
    }

    if len(message) > 500 {
      c.HTML(http.StatusOK, "contact-failure.html", gin.H{
        "email": email,
        "error": "Message is too long",
      })
      return
    }

    c.HTML(http.StatusOK, "contact-success.html", gin.H{
      "name": name,
      "email": email,
    })
  })
  r.Static("/templates", "./templates")
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}