package main

import (
	"embed"

	"github.com/aborroy/alf-opensearch/cmd"
	"github.com/aborroy/alf-opensearch/pkg"
)

//go:embed all:templates
var templateFs embed.FS

func main() {
	pkg.TemplateFs = templateFs
	cmd.TemplateFs = templateFs
	cmd.Execute()
}
