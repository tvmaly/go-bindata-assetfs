// example is a work in progress
// right now we want to ensure the dependencies are
// kept
package main

import (
	"fmt"
	"github.com/tmthrgd/chacha20"
	"github.com/tmthrgd/go-bindata"
	"github.com/tmthrgd/go-rand"
	"github.com/zach-klippenstein/goregen"
)

func main() {
	fmt.Println("force dep to include the imported dependencies")
}
