package main

import (
	"context"
	"os"

	"example.com/ee/lib/view"
	"example.com/ee/lib/yaml"
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
