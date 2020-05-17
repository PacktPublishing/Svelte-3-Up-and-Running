package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
)

// ErrorResponse is a struct used to respond to requests with an error message
type ErrorResponse struct {
	Error string `json:"error"`
}

// NewErrorResponse creates a new ErrorResponse object with the error message passed as argument
func NewErrorResponse(err string) ErrorResponse {
	return ErrorResponse{
		Error: err,
	}
}

// TokenClaims is a struct that extends jwt.StandardClaims and is used to define the list of claims in the id_token JWT
type TokenClaims struct {
	// Include the JWT standard claims, per RFC-7519
	jwt.StandardClaims

	// OpenID Connect standard claims (a subset of them)
	// See: https://openid.net/specs/openid-connect-core-1_0.html#Claims
	Name              string `json:"name"`
	GivenName         string `json:"given_name"`
	FamilyName        string `json:"family_name"`
	PreferredUsername string `json:"preferred_username"`
	Picture           string `json:"picture"`
	Email             string `json:"email"`
	EmailVerified     bool   `json:"email_verified"`
	Nonce             string `json:"nonce"`
}

// Cached signing keypair
var (
	cachedPubKey *rsa.PublicKey
	cachedPrvKey *rsa.PrivateKey
)

// SigningKeyPair loads (and generates if necessary) the RSA keypair used to sign and verify JWT tokens
func SigningKeyPair() (pub *rsa.PublicKey, prv *rsa.PrivateKey, err error) {
	// Respond from cache
	if cachedPubKey != nil && cachedPrvKey != nil {
		pub = cachedPubKey
		prv = cachedPrvKey
		return
	}

	// Load the private key from storage, if any
	prv, err = loadPrivateKey()
	if err != nil {
		return
	}

	// No key? Generate one
	if prv == nil {
		prv, err = genPrivateKey()
		if err != nil {
			return
		}
	}

	// Extract the public key
	pub = &prv.PublicKey

	// Cache response
	cachedPubKey = pub
	cachedPrvKey = prv

	return
}

// Loads the private key from storage, if any
func loadPrivateKey() (prv *rsa.PrivateKey, err error) {
	// Try requesting the private key from storage, if it exist
	var found bool
	keyData := bytes.Buffer{}
	found, _, err = storeInstance.Get("signing.key", &keyData)
	if err != nil || !found || keyData.Len() == 0 {
		return
	}

	// Decode the PEM-encoded key
	pemBlock, _ := pem.Decode(keyData.Bytes())
	if pemBlock.Type != "RSA PRIVATE KEY" {
		err = errors.New("Invalid PEM block type")
		return
	}
	prv, err = x509.ParsePKCS1PrivateKey(pemBlock.Bytes)

	return
}

// Generates a new RSA-2048 private key (and saves it in storage)
func genPrivateKey() (out *rsa.PrivateKey, err error) {
	// Generate a new key
	var prv *rsa.PrivateKey
	prv, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}

	// Encode to pem
	var pemBlock = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(prv),
	}
	pemBuf := bytes.Buffer{}
	err = pem.Encode(&pemBuf, pemBlock)
	if err != nil {
		return
	}

	// Store key
	_, err = storeInstance.Set("signing.key", &pemBuf, nil)
	if err != nil {
		return
	}

	// Return the private key
	out = prv
	return
}
