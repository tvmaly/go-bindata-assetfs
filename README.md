# go-bindata-assetfs

Serve embedded files from [tmthrgd/go-bindata](https://github.com/tmthrgd/go-bindata) with `net/http`.


[GoDoc](http://godoc.org/github.com/tvmaly/go-bindata-assetfs)

### Installation

Install with

    $ go get github.com/tmthrgd/go-bindata/...
    $ go get github.com/elazarl/go-bindata-assetfs/...

### Creating embedded data

Usage is identical to [tmthrgd/go-bindata](https://github.com/tmthrgd/go-bindata) usage,
instead of running `go-bindata` run `go-bindata-assetfs`.

The tool will create a `bindata_assetfs.go` file, which contains the embedded data.

A typical use case is

    $ go-bindata-assetfs data/...

### Using assetFS in your code

The generated file provides an `assetFS()` function that returns a `http.Filesystem`
wrapping the embedded files. What you usually want to do is:

    http.Handle("/", http.FileServer(assetFS()))

This would run an HTTP server serving the embedded files.

## Without running binary tool

You can always just run the `go-bindata` tool, and then

use

     import "github.com/tvmaly/go-bindata-assetfs"
     ...
     http.Handle("/",
        http.FileServer(
        &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: "data"}))

to serve files embedded from the `data` directory.

## Example of using it with [Gin](https://github.com/gin-gonic/gin) framework

 see [gin-contrib/static](https://github.com/gin-contrib/static)


## Background  history of this fork

There is some controversy around the repository github.com/jteeuwen.  The original username was deleted from github.  
Someone else ( not the original jtweeuwen owner ) created a new github account using jteeuwen and created a new repository go-bindata.

I had an older copy of the github.com/jteeuwen/go-bindata on my computer, but this was causing a number of errors on OSX.

I came across a fork ( github.com/tmthrgd/go-bindata ) of the original repository that was completely re-written and is well maintained.  

Parts of it were not compatible with github.com/elazarl/go-bindata-assetfs so this fork github.com/tvmaly/go-bindata-assetfs was created 
to be compatible with github.com/tmthrgd/go-bindata.

My primary motivation was have a way to embed the statis assets in the binary of my Gin based web projects.

This will work for other frameworks as it the http.FileServer interface
