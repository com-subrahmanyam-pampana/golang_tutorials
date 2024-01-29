// main.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAInLineHTMLPage() {
	var router *gin.Engine
	/*Set the router as the default one provided by Gin*/
	router = gin.Default()

	// Define the route for the index page and display the index.html template
	// To start with, we'll use an inline route handler. Later on, we'll create
	// standalone functions that will be used as route handlers.
	router.GET("/", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.Data(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"text/html; charset=utf-8",
			// Pass the data that the page uses (in this case, 'title')
			[]byte("<html> <body>Hello</body> </html>"),
		)

	})

	// Start serving the application
	router.Run()

}
