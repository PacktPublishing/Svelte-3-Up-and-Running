package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/PacktPublishing/Svelte.js-3-Proof-of-Concept/api-server/store"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
)

// Store singleton
var storeInstance store.Store

// Resources box
var resourcesBox *packr.Box

// Main method that launches the server
func main() {
	// Seed the pseudo-random number generator
	rand.Seed(time.Now().UnixNano())

	// Init the store instance
	storeAddress := os.Getenv("STORE_ADDRESS")
	if storeAddress == "" {
		storeAddress = "local:data"
	}
	var err error
	storeInstance, err = store.Get(storeAddress)
	if err != nil {
		panic(err)
	}

	// If we're in production mode, set Gin to "release" mode
	env := os.Getenv("ENV")
	if env != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Load the resources box
	resourcesBox = packr.New("resources", "./resources")

	// Start gin
	router := gin.Default()

	// Enable CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("authorization", "x-object-title", "x-object-date")
	corsConfig.AddExposeHeaders("date", "x-object-title", "x-object-date")
	router.Use(cors.New(corsConfig))

	// Add all routes
	router.GET("/", WelcomeHandler)
	router.GET("/authorize", AuthorizeHandler)
	router.GET("/jwks", JWKSHandler)
	router.POST("/authorize/token", AuthorizeTokenHandler)

	// Routes that require authorization
	{
		authorized := router.Group("/")
		authorized.Use(AuthMiddleware())
		authorized.GET("/object/:objectId", GetObjectHandler)
		authorized.POST("/object", PostObjectHandler)
		authorized.POST("/search", SearchHandler)
	}

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 4 << 20 // 4 MiB

	// Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "4343"
	} else {
		portNum, err := strconv.Atoi(port)
		if err != nil {
			panic(err)
		}
		port = strconv.Itoa(portNum)
	}

	// Address to bind to
	bind := os.Getenv("BIND")
	if bind == "" {
		bind = "127.0.0.1"
	}

	// HTTP Server
	server := &http.Server{
		Addr:           bind + ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Handle graceful shutdown on SIGINT
	idleConnsClosed := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)
		signal.Notify(s, os.Interrupt, syscall.SIGTERM)
		<-s

		// We received an interrupt signal, shut down.
		if err := server.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			fmt.Printf("HTTP server shutdown error: %v\n", err)
		}
		close(idleConnsClosed)
	}()

	// Start the server
	fmt.Printf("Starting server on http://%s\n", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		panic(err)
	}

	<-idleConnsClosed
}
