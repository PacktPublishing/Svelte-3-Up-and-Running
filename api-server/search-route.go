package main

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

type searchBody struct {
	TimeStart int64 `form:"start" json:"start"`
	TimeEnd   int64 `form:"end" json:"end"`
}

// SearchHandler is the handler for the POST /search endpoint, which searches for documents in the index
func SearchHandler(c *gin.Context) {
	// Get the clientId
	clientId := c.MustGet("clientId").(string)

	// Get the parameters from the body
	var body searchBody
	if err := c.ShouldBind(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Ensure the parameters are valid
	if body.TimeEnd < 0 || body.TimeStart < 0 || (body.TimeEnd > 0 && body.TimeEnd < body.TimeStart) {
		c.AbortWithStatusJSON(http.StatusBadRequest, NewErrorResponse("Invalid time range"))
		return
	}

	// Get the index
	index, _, err := getIndex(clientId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Iterate through the index and filter results we want
	n := 0
	for _, el := range index {
		if el.Date < body.TimeStart {
			continue
		}
		if body.TimeEnd > 0 && el.Date >= body.TimeEnd {
			continue
		}
		index[n] = el
		n++
	}
	index = index[:n]

	// Sort the results
	sort.Slice(index, func(i, j int) bool {
		return index[i].Date < index[j].Date
	})

	// Response
	c.JSON(http.StatusOK, index)
}
