# Svelte 3 Up and Running

<a href="https://www.packtpub.com/product/svelte-3-up-and-running/9781839213625"><img src="https://static.packt-cdn.com/products/9781839213625/cover/smaller" alt="Svelte 3 Up and Running" height="256px" align="right"></a>

This is the code repository for [Svelte 3 Up and Running](https://www.packtpub.com/product/svelte-3-up-and-running/9781839213625), published by Packt.

**A practical guide to building production-ready static web apps with Svelte 3**

## What is this book about?
Svelte is a modern JavaScript framework used to build static web apps that are fast and lean, as well as being fun for developers to use. This book is a concise and practical introduction for those who are new to the Svelte framework which will have you up to speed with building apps quickly, and teach you how to use Svelte 3 to build apps that offer a great app user experience (UX).

This book covers the following exciting features: 
* Understand why Svelte 3 is the go-to framework for building static web apps that offer great UX
* Explore the tool setup that makes it easier to build and debug Svelte apps
* Scaffold your web project and build apps using the Svelte framework
* Create Svelte components using the Svelte template syntax and its APIs
* Combine Svelte components to build apps that solve complex real-world problems

If you feel this book is for you, get your [copy](https://www.amazon.com/dp/1839213620) today!

<a href="https://www.packtpub.com/?utm_source=github&utm_medium=banner&utm_campaign=GitHubBanner"><img src="https://raw.githubusercontent.com/PacktPublishing/GitHub/master/GitHub.png" alt="https://www.packtpub.com/" border="5" /></a>

## Instructions and Navigations
All of the code is organized into folders. For example, Chapter02.

The code will look like the following:
```
if (test expression)
{
  Statement upon condition is true
}
```
### API server

Throughout the book, you'll be building a sample Journaling application with Svelte 3 that runs within a web browser. Like most front-end applications, the sample app comes with a back-end API server that is used to authenticate users and offer persistent storage for the data. 

### Pre-compiled binaries

The easiest way to launch the API server is to download a pre-compiled binary for your platform. Pre-compiled binaries are available in this GitHub repository in the [Releases](https://github.com/PacktPublishing/Svelte-3-Up-and-Running/releases/latest) section.

- Windows: [64-bit](https://github.com/PacktPublishing/Svelte-3-Up-and-Running/releases/download/v202008130606/api-server-v202008130606-win64.zip), [32-bit](https://github.com/PacktPublishing/Svelte-3-Up-and-Running/releases/download/v202008130606/api-server-v202008130606-win32.zip)
- macOS: [64-bit (Intel)](https://github.com/PacktPublishing/Svelte-3-Up-and-Running/releases/download/v202008130606/api-server-v202008130606-macos.tar.gz) (See note below)
- Linux: [amd64](https://github.com/PacktPublishing/Svelte-3-Up-and-Running/releases/download/v202008130606/api-server-v202008130606-linux-amd64.tar.gz), [386](https://github.com/PacktPublishing/Svelte-3-Up-and-Running/releases/download/v202008130606/api-server-v202008130606-linux-386.tar.gz), [armv7 (32-bit)](https://github.com/PacktPublishing/Svelte-3-Up-and-Running/releases/download/v202008130606/api-server-v202008130606-linux-armv7.tar.gz), [arm64](https://github.com/PacktPublishing/Svelte-3-Up-and-Running/releases/download/v202008130606/api-server-v202008130606-linux-arm64.tar.gz)

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

### Book's sample code

In the `ch*` folders, you can find the sample code for each chapter of the book.


**Following is what you need for this book:**
The book is for frontend or full-stack developers looking to build modern web apps with Svelte. Web developers with experience in leading frontend JavaScript frameworks who wish to learn Svelte will find this book useful. The book assumes a solid understanding of JavaScript and core HTML5 technologies. Basic understanding of modern frontend frameworks will be beneficial, but not necessary.

With the following software and hardware list you can run all code files present in the book (Chapter 1-7).

### Software and Hardware List

| Chapter  | Software required                   | OS required                        |
| -------- | ------------------------------------| -----------------------------------|
| 1 to 7        | Node.js 8 or higher with NPM                    | Windows, Mac OS X, and Linux (Any) |
| 1 to 7       | Webpack            | Windows, Mac OS X, and Linux (Any) |
| 1 to 7      | Svelte          | Windows, Mac OS X, and Linux (Any) |
|1 to 7      | Visual Studio Code           | Windows, Mac OS X, and Linux (Any) |
| 6    | Microsoft Azure           | Windows, Mac OS X, and Linux (Any) |

### Related products <Other books you may enjoy>
* Hands-On JavaScript High Performance [[Packt]](https://www.packtpub.com/product/hands-on-javascript-high-performance/9781838821098) [[Amazon]](https://www.amazon.com/dp/1788293770)

* Clean Code in JavaScript [[Packt]](https://www.packtpub.com/product/clean-code-in-javascript/9781789957648) [[Amazon]](https://www.amazon.com/dp/1788293770)

## Get to Know the Author
**Alessandro Segala**
is a Product Manager at Microsoft working on developer tools. He has over a decade of experience building full-stack web applications, having worked as a professional developer as well as contributing to multiple open source projects. Alessandro is the maintainer of svelte-spa-router, one of the most popular client-side routers for Svelte 3.

### Suggestions and Feedback
[Click here](https://docs.google.com/forms/d/e/1FAIpQLSdy7dATC6QmEL81FIUuymZ0Wy9vH1jHkvpY57OiMeKGqib_Ow/viewform) if you have any feedback or suggestions.
