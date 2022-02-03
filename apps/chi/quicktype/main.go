package main

import (
	"net/http"

	"github.com/go-chi/docgen"
)

func main() {
	http.ListenAndServe(":3000", routes())
}

func newMarkdownOpts() docgen.MarkdownOpts {
	return docgen.MarkdownOpts{
		ProjectPath:        "",
		Intro:              "",
		ForceRelativeLinks: false,
		URLMap:             map[string]string{},
	}
}
