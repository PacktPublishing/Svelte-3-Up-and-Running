package main

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// Maximum number of objects a client is allowed to store
const maxObjectsPerClient = 200

type postObjectBody struct {
	Content string `form:"content" json:"content"`
}

// PostObjectHandler is the handler for the POST /object endpoint, which stores a new object
func PostObjectHandler(c *gin.Context) {
	// Get the clientId
	clientId := c.MustGet("clientId").(string)

	// Check if we have a multipart form with a file
	var objectId string
	mpf, err := c.MultipartForm()
	if mpf != nil && len(mpf.File) > 0 {
		// Get the file from the input
		file, err := c.FormFile("file")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("Uploaded file 'file' is empty or invalid"))
			return
		}
		fp, err := file.Open()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer fp.Close()

		// Save the file
		objectId, err = storeFile(fp, clientId)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	} else {
		// Try binding the form and read the "content" value
		var body postObjectBody
		if err := c.ShouldBind(&body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		if body.Content == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("Field 'content' is empty"))
			return
		}

		// Save the data
		objectId, err = storeFile(bytes.NewBufferString(body.Content), clientId)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	// Update the index
	err = addToIndex(clientId, objectId)
	if err != nil {
		// Remove the file from storage
		_ = storeInstance.Delete(clientId+"/"+objectId, nil)

		// Respond with error
		if err.Error() == "client has stored too many objects" {
			c.AbortWithStatusJSON(http.StatusConflict, NewErrorResponse(err.Error()))
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	// Respond with the object ID
	c.Header("Location", "/object/"+objectId)
	c.Header("Content-Type", "application/octet-stream")
	c.JSON(201, map[string]string{"objectId": objectId})
}

func storeFile(fp io.Reader, clientId string) (objectId string, err error) {
	// Generate a new object Id
	var objectIdUUID uuid.UUID
	objectIdUUID, err = uuid.NewV4()
	if err != nil {
		return
	}
	objectId = objectIdUUID.String()

	// Store the file
	_, err = storeInstance.Set(clientId+"/"+objectId, fp, nil)

	return
}
