package app

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/andrenormanlang/common"
	"github.com/andrenormanlang/database"
	"github.com/andrenormanlang/views"
	"github.com/rs/zerolog/log"
)

func validateLink(link string) bool {
	for char := range link{
		char_val := int(char)
		is_uppercase := (char_val >= 65) && (char_val <= 90)
		is_lowercase := (char_val >= 97) && (char_val <= 122)
		is_sign := (char == '-') || (char == '_')

		if!(is_uppercase || is_lowercase || is_sign){
			return false
		}

	}
}


func pageHandler(c *gin.Context, app_settings common.AppSettings, database database.Database) ([]byte, error) {

	var page_binding common.PageLinkBinding

	err := c.ShouldBindUri(&page_binding)

	// Validate the page regex

	if err != nil || len(page_binding.Link) == 0 {
		// TODO : We should be serving error page
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page uri"})

		return nil, err
	}

	// Get the post with the ID
	post, err := database.GetPost(post_binding.Id)

	if err != nil || post.Content == "" {
		err = serveErrorPage(c, "given post not found", http.StatusNotFound, app_settings)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post Not Found"})
		}
		return nil, err
	}

	// Generate HTML page
	post.Content = string(mdToHTML([]byte(post.Content)))
	post_view := views.MakePostPage(post.Title, post.Content, app_settings.AppNavbar.Links)
	html_buffer := bytes.NewBuffer(nil)
	if err = post_view.Render(c, html_buffer); err != nil {
		log.Error().Msgf("could not render: %v", err)
	}

	return html_buffer.Bytes(), nil
}