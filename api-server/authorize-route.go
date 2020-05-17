package main

import (
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type authorizeQS struct {
	ClientId     string `form:"client_id"`
	Nonce        string `form:"nonce"`
	RedirectUri  string `form:"redirect_uri"`
	ResponseMode string `form:"response_mode"`
	ResponseType string `form:"response_type"`
	Scope        string `form:"scope"`
}

// AuthorizeHandler is the handler for the GET /authorize endpoint, which simulates an OAuth 2.0 "authorize" endpoint
func AuthorizeHandler(c *gin.Context) {
	// Get the parameters from the querystring
	var qs authorizeQS
	if err := c.ShouldBindQuery(&qs); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Ensure the required parameters are set in the Qs
	if qs.ClientId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("Parameter client_id is missing"))
		return
	}
	if qs.Nonce == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("Parameter nonce is missing"))
		return
	}
	if !strings.HasPrefix(qs.RedirectUri, "http://") && !strings.HasPrefix(qs.RedirectUri, "https://") {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("Parameter redirect_uri must be a URL beginning in 'http://' or 'https://'"))
		return
	}
	if qs.ResponseMode != "fragment" {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("Parameter response_mode must be 'fragment'"))
		return
	}
	if qs.ResponseType != "id_token" {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("Parameter response_type must be 'id_token'"))
		return
	}
	if qs.Scope != "openid profile" {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("Parameter scope must be 'openid profile' (with spaces encoded as %20)"))
		return
	}

	// Show the authorization form
	c.Header("Content-Type", "text/html")
	f, err := resourcesBox.Open("authorize-form.html")
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
