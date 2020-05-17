package main

import (
	"encoding/base64"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWKSHandler is the handler for the GET /jwks endpoint, which simulates the endpoint returning the JWK set
func JWKSHandler(c *gin.Context) {
	// Get the public key
	pub, _, err := SigningKeyPair()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Exponent must be 65537 (AQAB)
	if pub.E != 65537 {
		c.AbortWithError(http.StatusInternalServerError, errors.New("Invalid exponent: not 65537"))
	}

	// Generate the JWK set response
	n := base64.StdEncoding.EncodeToString(pub.N.Bytes())
	e := "AQAB"
	jwk := struct {
		Kty string `json:"kty"`
		Kid string `json:"kid"`
		Use string `json:"use"`
		N   string `json:"n"`
		E   string `json:"e"`
	}{
		Kty: "RSA",
		Kid: "1",
		Use: "sig",
		N:   n,
		E:   e,
	}
	response := struct {
		Keys []interface{} `json:"keys"`
	}{
		Keys: []interface{}{jwk},
	}

	c.JSON(http.StatusOK, response)
}
