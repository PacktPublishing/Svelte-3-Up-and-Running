package main

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Auth middleware that checks the Authorization header in the request
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check the Authorization header and remove spaces around the value (if any)
		auth := c.GetHeader("Authorization")
		auth = strings.TrimSpace(auth)

		// If the token begins with "Bearer ", remove that (we make it optional)
		if strings.HasPrefix(auth, "Bearer ") {
			auth = auth[7:]
		}

		if len(auth) != 0 {
			token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
				// Azure AD tokens are signed with RS256 method
				if token.Method.Alg() != "RS256" {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
				}

				// Get the signing key
				pub, _, err := SigningKeyPair()
				if err != nil {
					return nil, err
				}
				return pub, nil
			})
			if err == nil {
				// Check claims
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					// Perform some extra checks: valid iss, non-empty aud, then ensure exp and nbf are present (they were validated already)
					if claims["iss"] == "http://svelte-poc-server" &&
						claims["aud"] != "" &&
						claims["exp"] != "" &&
						claims["nbf"] != "" {
						// All good; store the claims in the current request's context
						c.Set("claims", claims)
						c.Set("clientId", claims["aud"])
						return
					}
				}
			}
		}

		// If we're still here, authentication has failed
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}
}
