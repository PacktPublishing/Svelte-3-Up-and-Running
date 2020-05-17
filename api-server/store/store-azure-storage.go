package store

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/Azure/azure-pipeline-go/pipeline"
	"github.com/Azure/azure-storage-blob-go/azblob"
)

// AzureStorage stores files on Azure Blob Storage
type AzureStorage struct {
	storageAccountName string
	storageContainer   string
	storagePipeline    pipeline.Pipeline
	storageURL         string
}

func (f *AzureStorage) Init(connection string) error {
	// Ensure the connection string is valid and extract the parts
	// connection mus start with "azureblob:" or "azure:"
	// Then it must contain the storage account container
	r := regexp.MustCompile("^(azureblob|azure):([a-z0-9][a-z0-9-]{2,62})$")
	match := r.FindStringSubmatch(connection)
	if match == nil || len(match) != 3 {
		return errors.New("invalid connection string for Azure Blob Storage")
	}
	f.storageContainer = match[2]

	// Get the storage account name and key from the environment
	name := os.Getenv("AZURE_STORAGE_ACCOUNT")
	key := os.Getenv("AZURE_STORAGE_ACCESS_KEY")
	if name == "" || key == "" {
		return errors.New("environmental variables AZURE_STORAGE_ACCOUNT and AZURE_STORAGE_ACCESS_KEY are not defined")
	}
	f.storageAccountName = name

	// Storage endpoint
	f.storageURL = fmt.Sprintf("https://%s.blob.core.windows.net/%s", f.storageAccountName, f.storageContainer)

	// Authenticate with Azure Storage
	credential, err := azblob.NewSharedKeyCredential(f.storageAccountName, key)
	if err != nil {
		return err
	}
	f.storagePipeline = azblob.NewPipeline(credential, azblob.PipelineOptions{
		Retry: azblob.RetryOptions{
			MaxTries: 3,
		},
	})

	return nil
}

func (f *AzureStorage) Get(name string, out io.Writer) (found bool, tag interface{}, err error) {
	if name == "" {
		err = errors.New("name is empty")
		return
	}

	found = true

	// Create the blob URL
	u, err := url.Parse(f.storageURL + "/" + name)
	if err != nil {
		return
	}
	blockBlobURL := azblob.NewBlockBlobURL(*u, f.storagePipeline)

	// Download the file
	resp, err := blockBlobURL.Download(context.TODO(), 0, azblob.CountToEnd, azblob.BlobAccessConditions{}, false)
	if err != nil {
		if stgErr, ok := err.(azblob.StorageError); !ok {
			err = fmt.Errorf("network error while downloading the file: %s", err.Error())
		} else {
			// Blob not found
			if stgErr.Response().StatusCode == http.StatusNotFound {
				err = nil
				found = false
				return
			}
			err = fmt.Errorf("azure Storage error while downloading the file: %s", stgErr.Response().Status)
		}
		return
	}
	body := resp.Body(azblob.RetryReaderOptions{
		MaxRetryRequests: 3,
	})
	defer body.Close()

	// Check if the file exists but it's empty
	if resp.ContentLength() == 0 {
		body.Close()
		found = false
		return
	}

	// Copy the response body to the out stream
	_, err = io.Copy(out, body)
	if err != nil {
		return
	}

	// Get the ETag
	tagObj := resp.ETag()
	tag = &tagObj

	return
}

func (f *AzureStorage) Set(name string, in io.Reader, tag interface{}) (tagOut interface{}, err error) {
	if name == "" {
		err = errors.New("name is empty")
		return
	}

	// Create the blob URL
	u, err := url.Parse(f.storageURL + "/" + name)
	if err != nil {
		return nil, err
	}
	blockBlobURL := azblob.NewBlockBlobURL(*u, f.storagePipeline)

	// If we have a tag (ETag), we will allow the upload to succeed only if the tag matches
	// If there's no ETag, then the upload can succeed only if there's no file already

	// Access conditions for blob uploads: disallow the operation if the blob already exists
	// See: https://docs.microsoft.com/en-us/rest/api/storageservices/specifying-conditional-headers-for-blob-service-operations#Subheading1
	var accessConditions azblob.BlobAccessConditions
	if tag == nil {
		// Uploads can succeed only if there's no blob at that path yet
		accessConditions = azblob.BlobAccessConditions{
			ModifiedAccessConditions: azblob.ModifiedAccessConditions{
				IfNoneMatch: "*",
			},
		}
	} else {
		// Uploads can succeed only if the file hasn't been modified since we downloaded it
		accessConditions = azblob.BlobAccessConditions{
			ModifiedAccessConditions: azblob.ModifiedAccessConditions{
				IfMatch: *tag.(*azblob.ETag),
			},
		}
	}

	resp, err := azblob.UploadStreamToBlockBlob(context.TODO(), in, blockBlobURL, azblob.UploadStreamToBlockBlobOptions{
		BufferSize:       3 * 1024 * 1024,
		MaxBuffers:       2,
		AccessConditions: accessConditions,
	})
	if err != nil {
		if stgErr, ok := err.(azblob.StorageError); !ok {
			return nil, fmt.Errorf("network error while uploading the file: %s", err.Error())
		} else {
			return nil, fmt.Errorf("Azure Storage error failed while uploading the file: %s", stgErr.Response().Status)
		}
	}

	// Get the ETag
	tagObj := resp.ETag()
	tagOut = &tagObj

	return tagOut, nil
}

func (f *AzureStorage) Delete(name string, tag interface{}) (err error) {
	if name == "" {
		err = errors.New("name is empty")
		return
	}

	// Create the blob URL
	u, err := url.Parse(f.storageURL + "/" + name)
	if err != nil {
		return
	}
	blockBlobURL := azblob.NewBlockBlobURL(*u, f.storagePipeline)

	// If we have a tag (ETag), we will allow the operation to succeed only if the tag matches
	// If there's no ETag, then it will always be allowed
	var accessConditions azblob.BlobAccessConditions
	if tag != nil {
		// Operation can succeed only if the file hasn't been modified since we downloaded it
		accessConditions = azblob.BlobAccessConditions{
			ModifiedAccessConditions: azblob.ModifiedAccessConditions{
				IfMatch: *tag.(*azblob.ETag),
			},
		}
	}

	// Delete the blob
	_, err = blockBlobURL.Delete(context.TODO(), azblob.DeleteSnapshotsOptionInclude, accessConditions)
	return
}
