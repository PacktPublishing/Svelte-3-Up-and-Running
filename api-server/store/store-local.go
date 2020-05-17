package store

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/PacktPublishing/Svelte.js-3-Proof-of-Concept/api-server/utils"

	homedir "github.com/mitchellh/go-homedir"
)

// Local is the local file system
// This implementation does not rely on tags, as it's assumed that concurrency isn't an issue on a single machine
type Local struct {
	basePath string
}

func (f *Local) Init(connection string) error {
	// Ensure that connection starts with "local:" or "file:"
	if !strings.HasPrefix(connection, "local:") && !strings.HasPrefix(connection, "file:") {
		return fmt.Errorf("invalid scheme")
	}

	// Get the path
	path := connection[strings.Index(connection, ":")+1:]

	// Expand the tilde if needed
	path, err := homedir.Expand(path)
	if err != nil {
		return err
	}

	// Get the absolute path
	path, err = filepath.Abs(path)
	if err != nil {
		return err
	}

	// Ensure the path ends with a /
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	// Lastly, ensure the path exists
	err = utils.EnsureFolder(path)
	if err != nil {
		return err
	}

	f.basePath = path

	return nil
}

func (f *Local) Get(name string, out io.Writer) (found bool, tag interface{}, err error) {
	if name == "" {
		err = errors.New("name is empty")
		return
	}

	found = true

	// Open the file
	file, err := os.Open(f.basePath + name)
	if err != nil {
		if os.IsNotExist(err) {
			found = false
			err = nil
		}
		return
	}
	defer file.Close()

	// Check if the file has any content
	stat, err := file.Stat()
	if err != nil {
		return
	}
	if stat.Size() == 0 {
		found = false
		return
	}

	// Copy the file to the out stream
	_, err = io.Copy(out, file)
	if err != nil {
		return
	}

	return
}

func (f *Local) Set(name string, in io.Reader, tag interface{}) (tagOut interface{}, err error) {
	if name == "" {
		err = errors.New("name is empty")
		return
	}

	// Create intermediate folders if needed
	dir := path.Dir(name)
	if dir != "" {
		err = os.MkdirAll(f.basePath+dir, os.ModePerm)
		if err != nil {
			return
		}
	}

	// Create the file
	file, err := os.Create(f.basePath + name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Write the stream to file
	_, err = io.Copy(file, in)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (f *Local) Delete(name string, tag interface{}) (err error) {
	if name == "" {
		err = errors.New("name is empty")
		return
	}

	// Delete the file
	err = os.Remove(f.basePath + name)
	return
}
