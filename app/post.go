package app

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/andrenormanlang/database"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/rs/zerolog/log"
)

type PostBinding struct {
	Id string `uri:"id" binding:"required"`
}

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	// render markdown to HTML
	return markdown.Render(doc, renderer)
}

func makePostHandler(database database.Database) func(*gin.Context) {
	return func(c *gin.Context) {

		// localhost:8080/post/{id} -extract the parameter from the url
		var post_binding PostBinding
		if err := c.ShouldBindUri(&post_binding); err != nil {
			// TODO redo this error handling
			c.JSON(400, gin.H{"msg": err})
			return
		}
		
		
		// Get the post with ID
		post_id, err := strconv.Atoi(post_binding.Id)
		if err != nil {
			// TODO redo this error handling
			c.JSON(400, gin.H{"msg": err})
			return
		}

		post, err := database.GetPost(post_id)
		if err != nil {
			// TODO redo this error handling
			c.JSON(400, gin.H{"msg": err})
			return
		}


		// Markdown to HTML the post content
		// It is easier to store the content in markdown format then convert it to HTML
		// ....
		post.Content = string(mdToHTML([]byte(post.Content)))
		
		
		// Serve the templated page
		log.Warn().Msgf("Post: %v\n", post)
		c.HTML(http.StatusOK, "post",gin.H{
			"Title": post.Title,
			"Content": template.HTML(post.Content),
		})
	}
}
