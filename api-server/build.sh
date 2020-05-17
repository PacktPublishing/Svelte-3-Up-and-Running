#!/bin/sh

set -e

# Delete existing builds if any
rm -rf bin || true
mkdir -p bin
BUILD_VERSION=$(date -u +'%Y%m%d%H%M%S')

# Build using Docker
GO_BUILDER_IMAGE=golang:1.14-alpine
docker run \
    --rm \
    -v "$PWD":/usr/src/myapp \
    -w /usr/src/myapp \
    --env BUILD_VERSION=$BUILD_VERSION \
    $GO_BUILDER_IMAGE \
    sh -c '
        set -e

        echo -e "###\nInstall the zip utility\n"
        apk add zip

        echo -e "\n###\nFetching modules\n"
        GO111MODULE=on \
        go get

        echo -e "\n###\nBuilding linux/amd64\n"
        # Disable CGO so the binary is fully static
        mkdir bin/api-server-v${BUILD_VERSION}-linux-amd64
        CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=amd64 \
        GO111MODULE=on \
        go build \
            -o bin/api-server-v${BUILD_VERSION}-linux-amd64/api-server
        cp LICENSE bin/api-server-v${BUILD_VERSION}-linux-amd64
        cp README.md bin/api-server-v${BUILD_VERSION}-linux-amd64
        (cd bin && tar -czvf api-server-v${BUILD_VERSION}-linux-amd64.tar.gz api-server-v${BUILD_VERSION}-linux-amd64)

        echo -e "\n###\nBuilding linux/386\n"
        mkdir bin/api-server-v${BUILD_VERSION}-linux-386
        CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=386 \
        GO111MODULE=on \
        go build \
            -o bin/api-server-v${BUILD_VERSION}-linux-386/api-server
        cp LICENSE bin/api-server-v${BUILD_VERSION}-linux-386
        cp README.md bin/api-server-v${BUILD_VERSION}-linux-386
        (cd bin && tar -czvf api-server-v${BUILD_VERSION}-linux-386.tar.gz api-server-v${BUILD_VERSION}-linux-386)

        echo -e "\n###\nBuilding linux/arm64\n"
        mkdir bin/api-server-v${BUILD_VERSION}-linux-arm64
        CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=arm64 \
        GO111MODULE=on \
        go build \
            -o bin/api-server-v${BUILD_VERSION}-linux-arm64/api-server
        cp LICENSE bin/api-server-v${BUILD_VERSION}-linux-arm64
        cp README.md bin/api-server-v${BUILD_VERSION}-linux-arm64
        (cd bin && tar -czvf api-server-v${BUILD_VERSION}-linux-arm64.tar.gz api-server-v${BUILD_VERSION}-linux-arm64)

        echo -e "\n###\nBuilding linux/arm\n"
        mkdir bin/api-server-v${BUILD_VERSION}-linux-armv7
        CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=arm \
        GOARM=7 \
        GO111MODULE=on \
        go build \
            -o bin/api-server-v${BUILD_VERSION}-linux-armv7/api-server
        cp LICENSE bin/api-server-v${BUILD_VERSION}-linux-armv7
        cp README.md bin/api-server-v${BUILD_VERSION}-linux-armv7
        (cd bin && tar -czvf api-server-v${BUILD_VERSION}-linux-armv7.tar.gz api-server-v${BUILD_VERSION}-linux-armv7)

        echo -e "\n###\nBuilding darwin/amd64\n"
        mkdir bin/api-server-v${BUILD_VERSION}-macos
        CGO_ENABLED=0 \
        GOOS=darwin \
        GOARCH=amd64 \
        GO111MODULE=on \
        go build \
            -o bin/api-server-v${BUILD_VERSION}-macos/api-server
        cp LICENSE bin/api-server-v${BUILD_VERSION}-macos
        cp README.md bin/api-server-v${BUILD_VERSION}-macos
        (cd bin && tar -czvf api-server-v${BUILD_VERSION}-macos.tar.gz api-server-v${BUILD_VERSION}-macos)

        echo -e "\n###\nBuilding windows/amd64\n"
        mkdir bin/api-server-v${BUILD_VERSION}-win64
        CGO_ENABLED=0 \
        GOOS=windows \
        GOARCH=amd64 \
        GO111MODULE=on \
        go build \
            -o bin/api-server-v${BUILD_VERSION}-win64/api-server.exe
        cp LICENSE bin/api-server-v${BUILD_VERSION}-win64
        cp README.md bin/api-server-v${BUILD_VERSION}-win64
        (cd bin && zip -r api-server-v${BUILD_VERSION}-win64.zip api-server-v${BUILD_VERSION}-win64)

        echo -e "\n###\nBuilding windows/386\n"
        mkdir bin/api-server-v${BUILD_VERSION}-win32
        CGO_ENABLED=0 \
        GOOS=windows \
        GOARCH=386 \
        GO111MODULE=on \
        go build \
            -o bin/api-server-v${BUILD_VERSION}-win32/api-server.exe
        cp LICENSE bin/api-server-v${BUILD_VERSION}-win32
        cp README.md bin/api-server-v${BUILD_VERSION}-win32
        (cd bin && zip -r api-server-v${BUILD_VERSION}-win32.zip api-server-v${BUILD_VERSION}-win32)
        '
echo -e"\n###\nCompilation done\n"
ls -al bin

# Build Docker container
docker build -t italypaleale/sveltebook:latest .
docker tag italypaleale/sveltebook:latest italypaleale/sveltebook:${BUILD_VERSION}
