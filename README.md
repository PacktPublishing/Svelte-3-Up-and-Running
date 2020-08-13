# Svelte 3 Up and Running, published by Packt

This repository contains the sample code for the [Svelte 3 Up and Running](https://www.packtpub.com/web-development/svelte-3-up-and-running) book by Alessandro Segala ([ItalyPaleAle](https://github.com/ItalyPaleAle)), published by Packt.

## API Server

Throughout the book, you'll be building a sample Journaling application with Svelte 3 that runs within a web browser. Like most front-end applications, the sample app comes with a back-end API server that is used to authenticate users and offer persistent storage for the data. 

### Pre-compiled binaries

The easiest way to launch the API server is to download a pre-compiled binary for your platform. Pre-compiled binaries are available in this GitHub repository in the [Releases](https://github.com/PacktPublishing/Svelte-3-Up-and-Running/releases/latest) section.

- Windows: [64-bit](https://github.com/PacktPublishing/Svelte-3-Up-and-Running/releases/download/v202008050636/api-server-v202008050636-win64.zip), [32-bit](https://github.com/PacktPublishing/Svelte-3-Up-and-Running/releases/download/v202008050636/api-server-v202008050636-win32.zip)
- macOS: [64-bit (Intel)](https://github.com/PacktPublishing/Svelte-3-Up-and-Running/releases/download/v202008050636/api-server-v202008050636-macos.tar.gz) (See note below)
- Linux: [amd64](https://github.com/PacktPublishing/Svelte-3-Up-and-Running/releases/download/v202008050636/api-server-v202008050636-linux-amd64.tar.gz), [386](https://github.com/PacktPublishing/Svelte-3-Up-and-Running/releases/download/v202008050636/api-server-v202008050636-linux-386.tar.gz), [armv7 (32-bit)](https://github.com/PacktPublishing/Svelte-3-Up-and-Running/releases/download/v202008050636/api-server-v202008050636-linux-armv7.tar.gz), [arm64](https://github.com/PacktPublishing/Svelte-3-Up-and-Running/releases/download/v202008050636/api-server-v202008050636-linux-arm64.tar.gz)

After downloading the archive for your platform and uncompressing it, launch the API server by double-clicking on the executable file. You should see the application running in a terminal window. You'll then be able to connect to the API server by opening [`http://localhost:4343`](http://localhost:4343) in your browser. To terminate the API server, close the terminal window.

> **Note for macOS users:**
>
> The pre-compiled binary is not signed with an Apple developer certificate, and Gatekeeper will refuse to run it in newer versions of macOS. If this happens to you, you will notice an error saying that the app is coming from an unidentified developer.
>
> To run the application on macOS  you can either (temporarily) disable Gatekeeper and allow unsigned applications (see the [Apple Support page](https://apple.co/2E3mVYP)) or run this command:
>
> ```
> xattr -rc path/to/application
> ````
>
> (where `path/to/application` is the location of the downloaded binary).

### Run with Docker

If you have Docker installed on your development machine, you can run the API server with:

```sh
docker run --rm -p 4343:4343 \
    -v ~/data:/data \
    italypaleale/sveltebook
```

Where `~/data` is the path on your local machine where the API server will store persistent data.

### Source code

Source code for the API server (written in Go) is present in the [api-server](/api-server) directory, and instructions to run from source are in that folder's README file.

## Book's sample code

In the `ch*` folders, you can find the sample code for each chapter of the book.
