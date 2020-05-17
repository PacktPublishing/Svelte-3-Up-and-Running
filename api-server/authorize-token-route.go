package main

import (
	"math/rand"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type authorizeTokenBody struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	ClientId string `form:"client_id" json:"client_id"`
	Nonce    string `form:"nonce" json:"nonce"`
}

// This is a sample app, so we're hardcoding the username and password people should use
// Of course, don't do this in a real-world app!!
const username = "svelte"
const password = "svelte"

// Token expiry in minutes
const tokenExpiry = 360 // 6 hours

// AuthorizeTokenHandler is the handler for the POST /authorize/token endpoint, which checks the supplied credentials and issues an id_token
func AuthorizeTokenHandler(c *gin.Context) {
	// Get the username and password
	var body authorizeTokenBody
	if err := c.ShouldBind(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Check the username and password
	// This is a sample app, so we're hardcoding the username and password people should use
	// Of course, don't do this in a real-world app!!
	if body.Username != username {
		c.AbortWithStatusJSON(http.StatusUnauthorized, NewErrorResponse("Invalid username or password"))
		return
	}
	if body.Password != password {
		c.AbortWithStatusJSON(http.StatusUnauthorized, NewErrorResponse("Invalid username or password"))
		return
	}
	if body.ClientId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("Parameter client_id is missing"))
		return
	}
	if body.Nonce == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("Parameter nonce is missing"))
		return
	}

	// Generate an id_token for a bogus user
	// Get the current time, which is also used as nbf (Not-Before), less 2 minutes to account for small clock skews
	now := time.Now()
	nbf := now.Add(-2 * time.Minute)
	picture := ""
	// Picture is 50% male and 50% female
	if rand.Uint32()%2 == 0 {
		picture = "https://italypalealedev.blob.core.windows.net/public/undraw_female_avatar_w3jk.png"
	} else {
		picture = "https://italypalealedev.blob.core.windows.net/public/undraw_male_avatar_323b.png"
	}
	claims := TokenClaims{
		// Standard JWT claims
		StandardClaims: jwt.StandardClaims{
			// Issuer: the address of this server (a placeholder since this can be self-hosted)
			Issuer: "http://svelte-poc-server",
			// Subject: the ID of the user
			Subject: "24e1717b-77e7-48ec-b9c0-bec695b489e2",
			// Audience: the client ID
			Audience: body.ClientId,
			// ExpiresAt: expires in 1 hour
			ExpiresAt: nbf.Add(time.Duration(tokenExpiry) * time.Hour).Unix(),
			// NotBefore: not valid before now
			NotBefore: nbf.Unix(),
			// IssuedAt: time the token was issued at
			IssuedAt: now.Unix(),
		},

		// Some claims related to OpenID Connect
		Name:              "Svelte User",
		GivenName:         "Svelte",
		FamilyName:        "User",
		PreferredUsername: "SvelteUser1",
		Picture:           picture,
		Email:             "svelteuser@example.com",
		EmailVerified:     true,
		Nonce:             body.Nonce,
	}

	// Generate the id_token JWT, and sign it with RS256 (uses public key cryptography, which is required with static web apps who can't safely store a symmetric key in their code)
	_, prvKey, err := SigningKeyPair()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kty"] = "RSA"
	token.Header["kid"] = "1"
	tokenString, err := token.SignedString(prvKey)

	// Return the JWT
	c.JSON(http.StatusOK, map[string]string{"id_token": tokenString})
}
