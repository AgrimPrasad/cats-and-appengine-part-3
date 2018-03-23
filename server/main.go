package main

import (
	"github.com/NYTimes/marvin"
	cats "github.com/jprobinson/cats-and-appengine-part-3"

	"google.golang.org/appengine"
)

func main() {
	marvin.Init(cats.NewService())
	appengine.Main()
}
