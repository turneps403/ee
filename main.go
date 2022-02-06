package main

import (
	"context"
	"os"

	"github.com/turneps403/ee/lib/view"
	"github.com/turneps403/ee/lib/yaml"
)

const confFile = ".ee.yaml"

func main() {
	initDir := ""
	if len(os.Args) > 1 {
		initDir = os.Args[1]
	}
	files := yaml.GetFiles(initDir, confFile)
	yamlObj := yaml.LoadAndMerge(files)
	ctx, cancel := context.WithCancel(context.Background())
	yamlTree := view.BuildTree(yamlObj, cancel)
	view.Run(yamlTree, ctx)
}
