package main

import (
	"os"

	"example.com/ee/lib/view"
	"example.com/ee/lib/yaml"
)

const confFile = ".ee.yaml"

func main() {
	// fmt.Println(exPath)
	initDir := ""
	if len(os.Args) > 1 {
		initDir = os.Args[1]
	}
	files := yaml.GetFiles(initDir, confFile)
	yamlObj := yaml.LoadAndMerge(files)
	yamlTree := view.BuildTree(yamlObj)
	view.Run(yamlTree)

	// 	config, _ := toml.Load(`
	// [postgres]
	// user = "pelletier"
	// password = "mypassword"`)
	// 	fmt.Println(config.Keys()[0])
	// 	return
	// 	// retrieve data directly
	// 	// user := config.Get("postgres.user").(string)

	// 	// // or using an intermediate object
	// 	// postgresConfig := config.Get("postgres").(*toml.Tree)
	// 	// password := postgresConfig.Get("password").(string)

	// 	rootDir := "."
	// 	root := tview.NewTreeNode(rootDir).SetColor(tcell.ColorRed)
	// 	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)

	// 	// A helper function which adds the files and directories of the given path
	// 	// to the given target node.
	// 	add := func(target *tview.TreeNode, path string) {
	// 		files, err := ioutil.ReadDir(path)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		for _, file := range files {
	// 			node := tview.NewTreeNode(file.Name()).
	// 				SetReference(filepath.Join(path, file.Name())).
	// 				// SetSelectable(file.IsDir())
	// 				SetSelectable(true)
	// 			if file.IsDir() {
	// 				node.SetColor(tcell.ColorGreen)
	// 			}
	// 			target.AddChild(node)
	// 		}
	// 	}

	// 	// Add the current directory to the root node.
	// 	add(root, rootDir)

	// 	// If a directory was selected, open it.
	// 	tree.SetSelectedFunc(func(node *tview.TreeNode) {
	// 		reference := node.GetReference()
	// 		if reference == nil {
	// 			return // Selecting the root node does nothing.
	// 		}
	// 		// fmt.Printf("Selected %s\n", reference.(string))
	// 		children := node.GetChildren()
	// 		if len(children) == 0 {
	// 			// Load and show files in this directory.
	// 			path := reference.(string)
	// 			add(node, path)
	// 		} else {
	// 			// Collapse if visible, expand if collapsed.
	// 			node.SetExpanded(!node.IsExpanded())
	// 		}
	// 	})

	// 	if err := tview.NewApplication().SetRoot(tree, true).Run(); err != nil {
	// 		panic(err)
	// 	}
}
