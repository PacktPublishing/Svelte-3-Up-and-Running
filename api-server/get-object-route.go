package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// GetObjectHandler is the handler for the GET /object/:objectId endpoint, which returns the contents of an object
func GetObjectHandler(c *gin.Context) {
	// Get the objectId
	objectId := c.Param("objectId")
	if objectId == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("empty objectId"))
		return
	}

	// Get the clientId
	clientId := c.MustGet("clientId").(string)

	// Ensure objectId is a UUID
	objectIdUUID, err := uuid.FromString(objectId)
	if err != nil || objectIdUUID.Version() != 4 {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Get the index to get the title
	index, _, err := getIndex(clientId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	title := ""
	var date time.Time
	found := false
	for _, el := range index {
		if el.ObjectId == objectId {
			found = true
			date = time.Unix(el.Date, 0)
			title = el.Title
			break
		}
	}
	if !found {
		// Object was not in the index
		c.AbortWithStatusJSON(http.StatusNotFound, NewErrorResponse("Object not found"))
		return
	}

	// Return the date and title as headers
	c.Header("x-object-date", date.In(time.FixedZone("GMT", 0)).Format(time.RFC1123))
	c.Header("x-object-title", title)

	// Get the object and return it to the client
	found, _, err = storeInstance.Get(clientId+"/"+objectId, c.Writer)
	if !found {
		c.AbortWithStatusJSON(http.StatusNotFound, NewErrorResponse("Object not found"))
		return
	}
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
