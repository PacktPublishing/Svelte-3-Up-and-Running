package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"time"
)

// IndexFile is the list of objects and dates for a clientId
type IndexFile []IndexElement

// IndexElement is an object in the IndexFile slice
type IndexElement struct {
	ObjectId string `json:"oid"`
	Date     int64  `json:"date"`
	Title    string `json:"title,omitempty"`
}

func getIndex(clientId string) (index IndexFile, tag interface{}, err error) {
	indexBuf := &bytes.Buffer{}
	found, tag, err := storeInstance.Get(clientId+"/_index.json", indexBuf)
	if found {
		// We have an index; read it
		read, err := ioutil.ReadAll(indexBuf)
		if err != nil {
			return index, nil, err
		}
		err = json.Unmarshal(read, &index)
		if err != nil {
			return index, nil, err
		}

		// Check if we're over the limit of objects
		if len(index) >= maxObjectsPerClient {
			return index, nil, errors.New("client has stored too many objects")
		}
	} else {
		// New index
		index = IndexFile{}
		tag = nil
	}

	return index, tag, nil
}

func addToIndex(clientId string, objectId string, title string) (err error) {
	// Get the current index
	index, tag, err := getIndex(clientId)
	if err != nil {
		return err
	}

	// Add the object
	index = append(index, IndexElement{
		ObjectId: objectId,
		Date:     time.Now().Unix(),
		Title:    title,
	})

	// Encode the updated index
	write, err := json.Marshal(index)
	if err != nil {
		return err
	}

	// Store the updated index
	_, err = storeInstance.Set(clientId+"/_index.json", bytes.NewBuffer(write), tag)
	if err != nil {
		return err
	}

	return nil
}
