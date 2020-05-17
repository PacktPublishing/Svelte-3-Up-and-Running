package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// WelcomeHandler is the handler for the GET / endpoint, which shows a welcome page
func WelcomeHandler(c *gin.Context) {
	// Show the welcome page
	c.Header("Content-Type", "text/html")
	f, err := resourcesBox.Open("welcome-page.html")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer f.Close()
	_, err = io.Copy(c.Writer, f)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
