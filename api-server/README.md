# API server

This folder contains the source code for the API server used by the sample app in the book.

The service includes common RESTful endpoints to store, retrieve, and search objects (journal entries).  It also features a mock OAuth 2.0 / OpenID Connect implementation to provide identity services. This was built from scratch, but it includes just the bare minimum features to support the needs of the sample front-end application.

## IMPORTANT NOTE

**Do not use this app or any of its code in production as-is.**

While this back-end service is functional, because its purpose is just to aid the development of the front-end application in the book, it is full of sub-optimal practices.  
This is especially key for the access management part, which is likely unsafe for any real-world application; instead, you should rely on your organization’s directory or, if building a consumer-facing app, on trustworthy identity service providers.

## Building the API server

The API server was written for **Go 1.14**, which needs to be installed before you can build the code.

You also need [packr v2](https://github.com/gobuffalo/packr/tree/master/v2) installed, which you can get with `go get -u github.com/gobuffalo/packr/v2/packr2`

1. Clone this repository (we use Go modules, so clone this outside of your GOPATH!): `git clone https://github.com/PacktPublishing/Svelte-3-Up-and-Running`
2. Navigate to the `api-server` folder within the cloned repository: `cd Svelte-3-Up-and-Running/api-server`
3. Run packr2: `packr2`
4. To run the application from source, run: `go run .`
5. To compile a binary for the platform you're currently using, run: `go build .`

## Building for all platforms

This folder contains a script that automatically builds the API server binaries for all supported platforms: `build.sh`.

This script requires Docker installed and running, and automatically builds binaries for Windows (64-bit and 32-bit), macOS (64-bit Intel), and Linux (amd64, 386, armv7, arm64). Additionally, the script builds a Docker image and tags it with `italypaleale/sveltebook:latest`.

To run the script:

```sh
packr2
./build.sh
```

## Runtime options

The API server supports a few runtime options that can be passed as environmental variables.

### Binding and port

You can configure the port and address the API server listens on with the following environmental variables:

- `PORT` sets the port to listen to (default is `4343`)
- `BIND` sets the address to bind to; the default is `127.0.0.1`, but to respond to requests from other nodes in the network you can set this to `0.0.0.0` (when running with Docker, the default is `0.0.0.0`)

### Persistent storage

By default, the API server stores its data (the journal entries) in a folder called `data` in the same directory as the binary. This can be configured with the **`STORE_ADDRESS`** environmental variable.

To store the data on the **local file system**, set `STORE_ADDRESS` to a string in the format: `local:folder_name`, where `folder_name` is the path to the folder where to store data, relative to the current directory. `local:data` is the default value for `STORE_ADDRESS`.

To use **Azure Blob Storage**:

1. Set the environmental variable `STORE_ADDRESS` to a string in the format: `azure:container_name`, where `container_name` is the name of the blob storage container in the Azure Storage account
2. Set the environmental variable `AZURE_STORAGE_ACCOUNT` with the name of your Azure Storage account
3. Set the environmental variable `AZURE_STORAGE_ACCESS_KEY` with the secret key of your Azure Storage account
