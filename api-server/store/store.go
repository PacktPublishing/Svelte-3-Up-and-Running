package store

import (
	"fmt"
	"io"
	"strings"
)

// Get returns a store for the given connection string
func Get(connection string) (store Store, err error) {
	store = nil

	// Get the name of the store
	pos := strings.Index(connection, ":")
	if pos < 1 {
		err = fmt.Errorf("invalid connection string")
		return
	}

	switch connection[0:pos] {
	case "file", "local":
		store = &Local{}
		err = store.Init(connection)
	case "azure", "azureblob":
		store = &AzureStorage{}
		err = store.Init(connection)
	default:
		err = fmt.Errorf("invalid connection string")
	}

	return
}

// Store is the interface for the store
type Store interface {
	// Init the object, by passing a connection string
	Init(connection string) error

	// Get returns a stream to a file in the filesystem
	// It also returns a tag (which might be empty) that should be passed to the Set method if you want to subsequentially update the contents of the file
	Get(name string, out io.Writer) (found bool, tag interface{}, err error)

	// Set writes a stream to the file in the filesystem
	// If you pass a tag, the implementation might use that to ensure that the file on the filesystem hasn't been changed since it was read (optional)
	Set(name string, in io.Reader, tag interface{}) (tagOut interface{}, err error)

	// Delete a file from the filesystem
	// If you pass a tag, the implementation might use that to ensure that the file on the filesystem hasn't been changed since it was read (optional)
	Delete(name string, tag interface{}) (err error)
}
